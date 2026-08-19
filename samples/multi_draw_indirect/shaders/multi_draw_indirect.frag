#version 460

/* Copyright (c) 2021-2026 Holochip Corporation
 * SPDX-License-Identifier: Apache-2.0
 */

layout(location = 1) in vec2 in_uv;
layout(location = 2) flat in uint in_texture_index;

layout(location = 0) out vec4 o_color;

layout(binding = 1, set = 0) uniform sampler2D textures[16];

void main(void)
{
	o_color = vec4(texture(textures[in_texture_index], in_uv));
	o_color.rgb *= 1.5;
}
