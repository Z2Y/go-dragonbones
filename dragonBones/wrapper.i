%module(directors="1", allprotected="1") wrapper

#define final

%{
#include "DragonBonesHeaders.h"
%}

%include "std_string.i"
%include "std_vector.i"
%include "wrapper_map.i"

%typemap(freearg) const char *rawData "";

/*%typemap(in) const char *rawData
%{
  printf("%llu\n", &$input);
  $1 = ($1_ltype)$input.p;
%}
*/

%rename(opLess) operator <;
%rename(opEqual) operator =;

%feature("director") Slot;
%feature("director") BaseFactory;
%feature("director") TextureData;
%feature("director") TextureAtlasData;
%feature("director") IArmatureProxy;

%include "DragonBones.h"
%include "BaseObject.h"
%include "Animation.h"
%include "TransformObject.h"
%include "IEventDispatcher.h"
%include "IArmatureProxy.h"
%include "IAnimatable.h"
%include "Slot.h"
%include "Armature.h"
%include "DragonBonesData.h"
%include "TextureAtlasData.h"
%include "ArmatureData.h"
%include "BaseFactory.h"
%include "Rectangle.h"
%include "Transform.h"
%include "Matrix.h"
%include "WorldClock.h"

%inline %{
template<class T>
std::size_t getTypeIndex(T*) {
    return typeid(T).hash_code();
}
%}

%extend dragonBones::WorldClock {
  void addArmature(dragonBones::Armature* value) {
    $self->add(value);
  }
}

%template(getSlotTypeIndex) getTypeIndex<SwigDirector_Slot>;

%template(getTextureAtlasDataTypeIndex) getTypeIndex<SwigDirector_TextureAtlasData>;

%template(borrowArmatureObject) dragonBones::BaseObject::borrowObject<dragonBones::Armature>;

%template(mapStringToTextureData) std::map< std::string,dragonBones::TextureData *,std::less< std::string > >;
%template(vectorString) std::vector<std::string>;
%template(vectorTextureData) std::vector<dragonBones::TextureData *>;
