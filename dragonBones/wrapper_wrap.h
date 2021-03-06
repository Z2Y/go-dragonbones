/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.2
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: .\wrapper.i

#ifndef SWIG_wrapper_WRAP_H_
#define SWIG_wrapper_WRAP_H_

class Swig_memory;

class SwigDirector_IArmatureProxy : public dragonBones::IArmatureProxy
{
 public:
  SwigDirector_IArmatureProxy(int swig_p);
  virtual ~SwigDirector_IArmatureProxy();
  virtual bool hasDBEventListener(std::string const &etype) const;
  virtual void dispatchDBEvent(std::string const &etype, dragonBones::EventObject *value);
  virtual void addDBEventListener(std::string const &etype, std::function< void (dragonBones::EventObject *) > const &listener);
  virtual void removeDBEventListener(std::string const &etype, std::function< void (dragonBones::EventObject *) > const &listener);
  virtual void dbInit(dragonBones::Armature *armature);
  virtual void dbClear();
  virtual void dbUpdate();
  virtual void dispose(bool disposeProxy);
  virtual dragonBones::Armature *getArmature() const;
  virtual dragonBones::Animation *getAnimation() const;
 private:
  intgo go_val;
  Swig_memory *swig_mem;
};

class SwigDirector_Slot : public dragonBones::Slot
{
 public:
  SwigDirector_Slot(int swig_p);
  virtual ~SwigDirector_Slot();
  void _swig_upcall__onClear() {
    dragonBones::Slot::_onClear();
  }
  virtual void _onClear();
  virtual std::size_t getClassTypeIndex() const;
  virtual void _initDisplay(void *value, bool isRetain);
  virtual void _disposeDisplay(void *value, bool isRelease);
  virtual void _onUpdateDisplay();
  virtual void _addDisplay();
  virtual void _replaceDisplay(void *value, bool isArmatureDisplay);
  virtual void _removeDisplay();
  virtual void _updateZOrder();
  virtual void _updateFrame();
  virtual void _updateMesh();
  virtual void _updateTransform();
  virtual void _identityTransform();
  virtual void _updateVisible();
  virtual void _updateBlendMode();
  virtual void _updateColor();
    using dragonBones::Slot::_displayDirty;
    using dragonBones::Slot::_zOrderDirty;
    using dragonBones::Slot::_visibleDirty;
    using dragonBones::Slot::_blendModeDirty;
    using dragonBones::Slot::_transformDirty;
    using dragonBones::Slot::_visible;
    using dragonBones::Slot::_displayIndex;
    using dragonBones::Slot::_animationDisplayIndex;
    using dragonBones::Slot::_cachedFrameIndex;
    using dragonBones::Slot::_localMatrix;
    using dragonBones::Slot::_displayDatas;
    using dragonBones::Slot::_displayList;
    using dragonBones::Slot::_rawDisplayDatas;
    using dragonBones::Slot::_boundingBoxData;
    using dragonBones::Slot::_textureData;
    using dragonBones::Slot::_display;
    using dragonBones::Slot::_childArmature;
    using dragonBones::Slot::_parent;
    using dragonBones::Slot::_getDefaultRawDisplayData;
    using dragonBones::Slot::_updateDisplay;
    using dragonBones::Slot::_updateDisplayData;
    using dragonBones::Slot::_updateGlobalTransformMatrix;
 private:
  intgo go_val;
  Swig_memory *swig_mem;
};

class SwigDirector_TextureAtlasData : public dragonBones::TextureAtlasData
{
 public:
  SwigDirector_TextureAtlasData(int swig_p);
  virtual ~SwigDirector_TextureAtlasData();
  void _swig_upcall__onClear() {
    dragonBones::TextureAtlasData::_onClear();
  }
  virtual void _onClear();
  virtual std::size_t getClassTypeIndex() const;
  virtual dragonBones::TextureData *createTexture() const;
  void _swig_upcall_addTexture(dragonBones::TextureData *value) {
    dragonBones::TextureAtlasData::addTexture(value);
  }
  virtual void addTexture(dragonBones::TextureData *value);
 private:
  intgo go_val;
  Swig_memory *swig_mem;
};

class SwigDirector_TextureData : public dragonBones::TextureData
{
 public:
  SwigDirector_TextureData(int swig_p);
  virtual ~SwigDirector_TextureData();
  void _swig_upcall__onClear() {
    dragonBones::TextureData::_onClear();
  }
  virtual void _onClear();
  virtual std::size_t getClassTypeIndex() const;
 private:
  intgo go_val;
  Swig_memory *swig_mem;
};

class SwigDirector_BaseFactory : public dragonBones::BaseFactory
{
 public:
  SwigDirector_BaseFactory(int swig_p, dragonBones::DataParser *dataParser);
  SwigDirector_BaseFactory(int swig_p);
  virtual ~SwigDirector_BaseFactory();
  bool _swig_upcall__isSupportMesh() const {
    return dragonBones::BaseFactory::_isSupportMesh();
  }
  virtual bool _isSupportMesh() const;
  dragonBones::TextureData *_swig_upcall__getTextureData(std::string const &textureAtlasName, std::string const &textureName) const {
    return dragonBones::BaseFactory::_getTextureData(textureAtlasName,textureName);
  }
  virtual dragonBones::TextureData *_getTextureData(std::string const &textureAtlasName, std::string const &textureName) const;
  bool _swig_upcall__fillBuildArmaturePackage(dragonBones::BuildArmaturePackage &dataPackage, std::string const &dragonBonesName, std::string const &armatureName, std::string const &skinName, std::string const &textureAtlasName) const {
    return dragonBones::BaseFactory::_fillBuildArmaturePackage(dataPackage,dragonBonesName,armatureName,skinName,textureAtlasName);
  }
  virtual bool _fillBuildArmaturePackage(dragonBones::BuildArmaturePackage &dataPackage, std::string const &dragonBonesName, std::string const &armatureName, std::string const &skinName, std::string const &textureAtlasName) const;
  void _swig_upcall__buildBones(dragonBones::BuildArmaturePackage const &dataPackage, dragonBones::Armature *armature) const {
    dragonBones::BaseFactory::_buildBones(dataPackage,armature);
  }
  virtual void _buildBones(dragonBones::BuildArmaturePackage const &dataPackage, dragonBones::Armature *armature) const;
  void _swig_upcall__buildSlots(dragonBones::BuildArmaturePackage const &dataPackage, dragonBones::Armature *armature) const {
    dragonBones::BaseFactory::_buildSlots(dataPackage,armature);
  }
  virtual void _buildSlots(dragonBones::BuildArmaturePackage const &dataPackage, dragonBones::Armature *armature) const;
  dragonBones::Armature *_swig_upcall__buildChildArmature(dragonBones::BuildArmaturePackage const *dataPackage, dragonBones::Slot *slot, dragonBones::DisplayData *displayData) const {
    return dragonBones::BaseFactory::_buildChildArmature(dataPackage,slot,displayData);
  }
  virtual dragonBones::Armature *_buildChildArmature(dragonBones::BuildArmaturePackage const *dataPackage, dragonBones::Slot *slot, dragonBones::DisplayData *displayData) const;
  std::pair< void *, dragonBones::DisplayType > _swig_upcall__getSlotDisplay(dragonBones::BuildArmaturePackage const *dataPackage, dragonBones::DisplayData *displayData, dragonBones::DisplayData *rawDisplayData, dragonBones::Slot *slot) const {
    return dragonBones::BaseFactory::_getSlotDisplay(dataPackage,displayData,rawDisplayData,slot);
  }
  virtual std::pair< void *, dragonBones::DisplayType > _getSlotDisplay(dragonBones::BuildArmaturePackage const *dataPackage, dragonBones::DisplayData *displayData, dragonBones::DisplayData *rawDisplayData, dragonBones::Slot *slot) const;
  virtual dragonBones::TextureAtlasData *_buildTextureAtlasData(dragonBones::TextureAtlasData *textureAtlasData, void *textureAtlas) const;
  virtual dragonBones::Armature *_buildArmature(dragonBones::BuildArmaturePackage const &dataPackage) const;
  virtual dragonBones::Slot *_buildSlot(dragonBones::BuildArmaturePackage const &dataPackage, dragonBones::SlotData const *slotData, dragonBones::Armature *armature) const;
  dragonBones::DragonBonesData *_swig_upcall_parseDragonBonesData__SWIG_0(char const *rawData, std::string const &name, float scale) {
    return dragonBones::BaseFactory::parseDragonBonesData(rawData,name,scale);
  }
  virtual dragonBones::DragonBonesData *parseDragonBonesData(char const *rawData, std::string const &name, float scale);
  dragonBones::DragonBonesData *_swig_upcall_parseDragonBonesData__SWIG_1(char const *rawData, std::string const &name) {
    return dragonBones::BaseFactory::parseDragonBonesData(rawData,name);
  }
  virtual dragonBones::DragonBonesData *parseDragonBonesData(char const *rawData, std::string const &name);
  dragonBones::DragonBonesData *_swig_upcall_parseDragonBonesData__SWIG_2(char const *rawData) {
    return dragonBones::BaseFactory::parseDragonBonesData(rawData);
  }
  virtual dragonBones::DragonBonesData *parseDragonBonesData(char const *rawData);
  dragonBones::TextureAtlasData *_swig_upcall_parseTextureAtlasData__SWIG_0(char const *rawData, void *textureAtlas, std::string const &name, float scale) {
    return dragonBones::BaseFactory::parseTextureAtlasData(rawData,textureAtlas,name,scale);
  }
  virtual dragonBones::TextureAtlasData *parseTextureAtlasData(char const *rawData, void *textureAtlas, std::string const &name, float scale);
  dragonBones::TextureAtlasData *_swig_upcall_parseTextureAtlasData__SWIG_1(char const *rawData, void *textureAtlas, std::string const &name) {
    return dragonBones::BaseFactory::parseTextureAtlasData(rawData,textureAtlas,name);
  }
  virtual dragonBones::TextureAtlasData *parseTextureAtlasData(char const *rawData, void *textureAtlas, std::string const &name);
  dragonBones::TextureAtlasData *_swig_upcall_parseTextureAtlasData__SWIG_2(char const *rawData, void *textureAtlas) {
    return dragonBones::BaseFactory::parseTextureAtlasData(rawData,textureAtlas);
  }
  virtual dragonBones::TextureAtlasData *parseTextureAtlasData(char const *rawData, void *textureAtlas);
  void _swig_upcall_addDragonBonesData__SWIG_0(dragonBones::DragonBonesData *data, std::string const &name) {
    dragonBones::BaseFactory::addDragonBonesData(data,name);
  }
  virtual void addDragonBonesData(dragonBones::DragonBonesData *data, std::string const &name);
  void _swig_upcall_addDragonBonesData__SWIG_1(dragonBones::DragonBonesData *data) {
    dragonBones::BaseFactory::addDragonBonesData(data);
  }
  virtual void addDragonBonesData(dragonBones::DragonBonesData *data);
  void _swig_upcall_removeDragonBonesData__SWIG_0(std::string const &name, bool disposeData) {
    dragonBones::BaseFactory::removeDragonBonesData(name,disposeData);
  }
  virtual void removeDragonBonesData(std::string const &name, bool disposeData);
  void _swig_upcall_removeDragonBonesData__SWIG_1(std::string const &name) {
    dragonBones::BaseFactory::removeDragonBonesData(name);
  }
  virtual void removeDragonBonesData(std::string const &name);
  void _swig_upcall_addTextureAtlasData__SWIG_0(dragonBones::TextureAtlasData *data, std::string const &name) {
    dragonBones::BaseFactory::addTextureAtlasData(data,name);
  }
  virtual void addTextureAtlasData(dragonBones::TextureAtlasData *data, std::string const &name);
  void _swig_upcall_addTextureAtlasData__SWIG_1(dragonBones::TextureAtlasData *data) {
    dragonBones::BaseFactory::addTextureAtlasData(data);
  }
  virtual void addTextureAtlasData(dragonBones::TextureAtlasData *data);
  void _swig_upcall_removeTextureAtlasData__SWIG_0(std::string const &name, bool disposeData) {
    dragonBones::BaseFactory::removeTextureAtlasData(name,disposeData);
  }
  virtual void removeTextureAtlasData(std::string const &name, bool disposeData);
  void _swig_upcall_removeTextureAtlasData__SWIG_1(std::string const &name) {
    dragonBones::BaseFactory::removeTextureAtlasData(name);
  }
  virtual void removeTextureAtlasData(std::string const &name);
  dragonBones::ArmatureData *_swig_upcall_getArmatureData__SWIG_0(std::string const &name, std::string const &dragonBonesName) const {
    return dragonBones::BaseFactory::getArmatureData(name,dragonBonesName);
  }
  virtual dragonBones::ArmatureData *getArmatureData(std::string const &name, std::string const &dragonBonesName) const;
  dragonBones::ArmatureData *_swig_upcall_getArmatureData__SWIG_1(std::string const &name) const {
    return dragonBones::BaseFactory::getArmatureData(name);
  }
  virtual dragonBones::ArmatureData *getArmatureData(std::string const &name) const;
  void _swig_upcall_clear__SWIG_0(bool disposeData) {
    dragonBones::BaseFactory::clear(disposeData);
  }
  virtual void clear(bool disposeData);
  void _swig_upcall_clear__SWIG_1() {
    dragonBones::BaseFactory::clear();
  }
  virtual void clear();
  dragonBones::Armature *_swig_upcall_buildArmature__SWIG_0(std::string const &armatureName, std::string const &dragonBonesName, std::string const &skinName, std::string const &textureAtlasName) const {
    return dragonBones::BaseFactory::buildArmature(armatureName,dragonBonesName,skinName,textureAtlasName);
  }
  virtual dragonBones::Armature *buildArmature(std::string const &armatureName, std::string const &dragonBonesName, std::string const &skinName, std::string const &textureAtlasName) const;
  dragonBones::Armature *_swig_upcall_buildArmature__SWIG_1(std::string const &armatureName, std::string const &dragonBonesName, std::string const &skinName) const {
    return dragonBones::BaseFactory::buildArmature(armatureName,dragonBonesName,skinName);
  }
  virtual dragonBones::Armature *buildArmature(std::string const &armatureName, std::string const &dragonBonesName, std::string const &skinName) const;
  dragonBones::Armature *_swig_upcall_buildArmature__SWIG_2(std::string const &armatureName, std::string const &dragonBonesName) const {
    return dragonBones::BaseFactory::buildArmature(armatureName,dragonBonesName);
  }
  virtual dragonBones::Armature *buildArmature(std::string const &armatureName, std::string const &dragonBonesName) const;
  dragonBones::Armature *_swig_upcall_buildArmature__SWIG_3(std::string const &armatureName) const {
    return dragonBones::BaseFactory::buildArmature(armatureName);
  }
  virtual dragonBones::Armature *buildArmature(std::string const &armatureName) const;
  void _swig_upcall_replaceDisplay(dragonBones::Slot *slot, dragonBones::DisplayData *displayData, int displayIndex) const {
    dragonBones::BaseFactory::replaceDisplay(slot,displayData,displayIndex);
  }
  virtual void replaceDisplay(dragonBones::Slot *slot, dragonBones::DisplayData *displayData, int displayIndex) const;
  bool _swig_upcall_replaceSlotDisplay__SWIG_0(std::string const &dragonBonesName, std::string const &armatureName, std::string const &slotName, std::string const &displayName, dragonBones::Slot *slot, int displayIndex) const {
    return dragonBones::BaseFactory::replaceSlotDisplay(dragonBonesName,armatureName,slotName,displayName,slot,displayIndex);
  }
  virtual bool replaceSlotDisplay(std::string const &dragonBonesName, std::string const &armatureName, std::string const &slotName, std::string const &displayName, dragonBones::Slot *slot, int displayIndex) const;
  bool _swig_upcall_replaceSlotDisplay__SWIG_1(std::string const &dragonBonesName, std::string const &armatureName, std::string const &slotName, std::string const &displayName, dragonBones::Slot *slot) const {
    return dragonBones::BaseFactory::replaceSlotDisplay(dragonBonesName,armatureName,slotName,displayName,slot);
  }
  virtual bool replaceSlotDisplay(std::string const &dragonBonesName, std::string const &armatureName, std::string const &slotName, std::string const &displayName, dragonBones::Slot *slot) const;
  bool _swig_upcall_replaceSlotDisplayList(std::string const &dragonBonesName, std::string const &armatureName, std::string const &slotName, dragonBones::Slot *slot) const {
    return dragonBones::BaseFactory::replaceSlotDisplayList(dragonBonesName,armatureName,slotName,slot);
  }
  virtual bool replaceSlotDisplayList(std::string const &dragonBonesName, std::string const &armatureName, std::string const &slotName, dragonBones::Slot *slot) const;
  bool _swig_upcall_replaceSkin__SWIG_0(dragonBones::Armature *armature, dragonBones::SkinData *skin, bool isOverride, std::vector< std::string > const *exclude) const {
    return dragonBones::BaseFactory::replaceSkin(armature,skin,isOverride,exclude);
  }
  virtual bool replaceSkin(dragonBones::Armature *armature, dragonBones::SkinData *skin, bool isOverride, std::vector< std::string > const *exclude) const;
  bool _swig_upcall_replaceSkin__SWIG_1(dragonBones::Armature *armature, dragonBones::SkinData *skin, bool isOverride) const {
    return dragonBones::BaseFactory::replaceSkin(armature,skin,isOverride);
  }
  virtual bool replaceSkin(dragonBones::Armature *armature, dragonBones::SkinData *skin, bool isOverride) const;
  bool _swig_upcall_replaceSkin__SWIG_2(dragonBones::Armature *armature, dragonBones::SkinData *skin) const {
    return dragonBones::BaseFactory::replaceSkin(armature,skin);
  }
  virtual bool replaceSkin(dragonBones::Armature *armature, dragonBones::SkinData *skin) const;
  bool _swig_upcall_replaceAnimation__SWIG_0(dragonBones::Armature *armature, dragonBones::ArmatureData *armatureData, bool isReplaceAll) const {
    return dragonBones::BaseFactory::replaceAnimation(armature,armatureData,isReplaceAll);
  }
  virtual bool replaceAnimation(dragonBones::Armature *armature, dragonBones::ArmatureData *armatureData, bool isReplaceAll) const;
  bool _swig_upcall_replaceAnimation__SWIG_1(dragonBones::Armature *armature, dragonBones::ArmatureData *armatureData) const {
    return dragonBones::BaseFactory::replaceAnimation(armature,armatureData);
  }
  virtual bool replaceAnimation(dragonBones::Armature *armature, dragonBones::ArmatureData *armatureData) const;
    using dragonBones::BaseFactory::_jsonParser;
    using dragonBones::BaseFactory::_binaryParser;
    using dragonBones::BaseFactory::_dragonBonesDataMap;
    using dragonBones::BaseFactory::_textureAtlasDataMap;
    using dragonBones::BaseFactory::_dragonBones;
    using dragonBones::BaseFactory::_dataParser;
 private:
  intgo go_val;
  Swig_memory *swig_mem;
};

#endif
