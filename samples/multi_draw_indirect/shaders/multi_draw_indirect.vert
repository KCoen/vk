#version 460

/* Copyright (c) 2021-2026 Holochip Corporation
 * SPDX-License-Identifier: Apache-2.0
 */

layout(location = 0) in vec3 position;
layout(location = 1) in vec2 uv;
layout(location = 2) in vec3 sphere_center;
layout(location = 3) in float sphere_radius;
layout(location = 4) in uint texture_index;

layout(set = 0, binding = 2) uniform GlobalUniform
{
	mat4 view;
	mat4 proj;
	mat4 proj_view;
	uint model_count;
}
global_uniform;

layout(location = 1) out vec2 o_uv;
layout(location = 2) out uint o_texture_index;

void main(void)
{
	o_uv            = uv;
	gl_Position     = global_uniform.proj * global_uniform.view * vec4(position, 1.0);
	o_texture_index = texture_index;
}
