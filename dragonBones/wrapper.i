%module(directors="1") wrapper

%{
#include "DragonBonesHeaders.h"
%}

%rename(opLess) operator <;

%feature("director") Slot;
%feature("director") BaseFactory;
%feature("director") TextureData;
%feature("director") TextureAtlasData;

%include "std_string.i"

%include "DragonBones.h"
%include "BaseObject.h"
%include "BaseFactory.h"
%include "TransformObject.h"
%include "IEventDispatcher.h"
%include "IArmatureProxy.h"
%include "Slot.h"
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
