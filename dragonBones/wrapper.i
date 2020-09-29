%module(directors="1") wrapper

#define final

%{
#include "DragonBonesHeaders.h"
%}

%rename(opLess) operator <;

%feature("director") Slot;
%feature("director") BaseFactory;
%feature("director") TextureData;
%feature("director") TextureAtlasData;
%feature("director") IArmatureProxy;

%include "std_string.i"

%include "DragonBones.h"
%include "BaseObject.h"
%include "Animation.h"
%include "BaseFactory.h"
%include "TransformObject.h"
%include "IEventDispatcher.h"
%include "IArmatureProxy.h"
%include "IAnimatable.h"
%include "Slot.h"
%include "Armature.h"
%include "DragonBonesData.h"
%include "TextureAtlasData.h"
%include "DataParser.h"
%include "JSONDataParser.h"

%inline %{
template<class T>
std::size_t getTypeIndex(T*) {
    return typeid(T).hash_code();
}
%}

%template(getSlotTypeIndex) getTypeIndex<SwigDirector_Slot>;

%template(getTextureAtlasDataTypeIndex) getTypeIndex<SwigDirector_TextureAtlasData>;

%template(borrowArmatureObject) dragonBones::BaseObject::borrowObject<dragonBones::Armature>;
