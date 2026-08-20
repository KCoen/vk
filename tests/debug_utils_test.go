package tests

import (
	"runtime"
	"strings"
	"sync/atomic"
	"testing"
	"unsafe"

	"github.com/KCoen/vk/extensions/ext_debug_utils"
	"github.com/KCoen/vk/vulkan"
)

func TestDebugUtilsCallback(t *testing.T) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if err := vulkan.Init(); err != nil {
		t.Skipf("Vulkan loader not available on host: %v", err)
	}

	// 1. Enumerate available extensions to ensure VK_EXT_debug_utils is supported
	exts, res := vulkan.EnumerateInstanceExtensionProperties("")
	if res != vulkan.SUCCESS {
		t.Fatalf("Failed to enumerate instance extensions: %v", res)
	}

	debugUtilsSupported := false
	for _, ext := range exts {
		name := vulkan.ByteSliceToString(ext.ExtensionName[:])
		if name == ext_debug_utils.ExtensionName {
			debugUtilsSupported = true
			break
		}
	}

	if !debugUtilsSupported {
		t.Skipf("Extension %s not supported by Vulkan driver/loader on this host", ext_debug_utils.ExtensionName)
	}

	// 2. Check for Validation layer availability
	layerProps, _ := vulkan.EnumerateInstanceLayerProperties()
	var enabledLayers []string
	for _, lp := range layerProps {
		layerName := vulkan.ByteSliceToString(lp.LayerName[:])
		if layerName == "VK_LAYER_KHRONOS_validation" {
			enabledLayers = append(enabledLayers, layerName)
			t.Logf("Found validation layer: %s", layerName)
			break
		}
	}

	// 3. Setup debug callback
	var callbackInvocations int32
	var lastReceivedMessage string

	debugCallback := func(
		messageSeverity vulkan.DebugUtilsMessageSeverityFlagBitsEXT,
		messageTypes vulkan.DebugUtilsMessageTypeFlagsEXT,
		callbackData *vulkan.RawDebugUtilsMessengerCallbackDataEXT,
		userData unsafe.Pointer,
	) vulkan.Bool32 {
		atomic.AddInt32(&callbackInvocations, 1)
		if callbackData != nil && callbackData.Message != nil {
			msg := vulkan.NullTerminatedToString(callbackData.Message)
			lastReceivedMessage = msg
		}
		return vulkan.False
	}

	messengerInfo := vulkan.NewDebugUtilsMessengerCreateInfoEXT()
	messengerInfo.MessageSeverity = vulkan.DEBUG_UTILS_MESSAGE_SEVERITY_VERBOSE_BIT_EXT |
		vulkan.DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT |
		vulkan.DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT |
		vulkan.DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT
	messengerInfo.MessageType = vulkan.DEBUG_UTILS_MESSAGE_TYPE_GENERAL_BIT_EXT |
		vulkan.DEBUG_UTILS_MESSAGE_TYPE_VALIDATION_BIT_EXT |
		vulkan.DEBUG_UTILS_MESSAGE_TYPE_PERFORMANCE_BIT_EXT
	// Directly assign Go function without manual purego.NewCallback!
	messengerInfo.UserCallback = debugCallback

	// 4. Create instance and chain messengerInfo via SetNext()
	appInfo := vulkan.ApplicationInfo{
		ApplicationName:    "DebugCallbackTest",
		ApplicationVersion: vulkan.MakeVersion(1, 0, 0),
		EngineName:         "vk_google",
		EngineVersion:      vulkan.MakeVersion(1, 0, 0),
		ApiVersion:         vulkan.MakeVersion(1, 0, 0),
	}

	instanceCreateInfo := vulkan.InstanceCreateInfo{
		ApplicationInfo:       &appInfo,
		EnabledExtensionNames: []string{ext_debug_utils.ExtensionName},
		EnabledLayerNames:     enabledLayers,
	}

	// Use the generated SetNext method to chain the debug messenger create info
	instanceCreateInfo.SetNext(messengerInfo)

	instance, res := vulkan.CreateInstance(&instanceCreateInfo, nil)
	if res != vulkan.SUCCESS {
		t.Fatalf("CreateInstance failed: %v", res)
	}
	defer vulkan.DestroyInstance(instance, nil)

	// Initialize branch and extension procedure addresses
	vulkan.InitCommands(instance, 0)
	ext_debug_utils.Init(instance, 0)

	// 5. Create explicit debug messenger
	messenger, res := ext_debug_utils.CreateDebugUtilsMessengerEXT(instance, messengerInfo, nil)
	if res != vulkan.SUCCESS {
		t.Fatalf("CreateDebugUtilsMessengerEXT failed: %v", res)
	}
	defer ext_debug_utils.DestroyDebugUtilsMessengerEXT(instance, messenger, nil)

	// 6. Submit custom debug message to trigger the callback
	testMessage := "Antigravity pure-Go Vulkan debug verification message"
	callbackData := vulkan.DebugUtilsMessengerCallbackDataEXT{
		Message: testMessage,
	}

	ext_debug_utils.SubmitDebugUtilsMessageEXT(
		instance,
		vulkan.DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT,
		vulkan.DEBUG_UTILS_MESSAGE_TYPE_GENERAL_BIT_EXT,
		&callbackData,
	)

	// 7. Validate that the callback was invoked and received the message
	invocations := atomic.LoadInt32(&callbackInvocations)
	if invocations == 0 {
		t.Fatalf("Debug callback was not triggered!")
	}

	if !strings.Contains(lastReceivedMessage, testMessage) {
		t.Fatalf("Expected message containing %q, got %q", testMessage, lastReceivedMessage)
	}

	t.Logf("Successfully verified debug callback! Invocations: %d, Message: %q", invocations, lastReceivedMessage)
}
