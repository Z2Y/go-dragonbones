/**
 * The MIT License (MIT)
 *
 * Copyright (c) 2012-2018 DragonBones team and other contributors
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 * 
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 * 
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */
#ifndef DRAGONBONES_HEADERS_H
#define DRAGONBONES_HEADERS_H

// core
#include "DragonBones.h"
#include "BaseObject.h"

// geom
#include "Matrix.h"
#include "Transform.h"
#include "ColorTransform.h"
#include "Point.h"
#include "Rectangle.h"

// model
#include "TextureAtlasData.h"
#include "UserData.h"
#include "DragonBonesData.h"
#include "ArmatureData.h"
#include "ConstraintData.h"
#include "CanvasData.h"
#include "SkinData.h"
#include "DisplayData.h"
#include "BoundingBoxData.h"
#include "AnimationData.h"
#include "AnimationConfig.h"

// armature
#include "IArmatureProxy.h"
#include "Armature.h"
#include "TransformObject.h"
#include "Bone.h"
#include "Slot.h"
#include "Constraint.h"
#include "DeformVertices.h"

// animation
#include "IAnimatable.h"
#include "WorldClock.h"
#include "Animation.h"
#include "AnimationState.h"
#include "BaseTimelineState.h"
#include "TimelineState.h"

// event
#include "EventObject.h"
#include "IEventDispatcher.h"

#ifndef EGRET_WASM

// parser
#include "DataParser.h"
#include "JSONDataParser.h"
#include "BinaryDataParser.h"

// factory
#include "BaseFactory.h"
#endif // EGRET_WASM

#endif // DRAGONBONES_HEADERS_H
