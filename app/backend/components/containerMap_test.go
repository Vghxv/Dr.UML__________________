package components

import (
	"testing"

	"Dr.uml/backend/mocks"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_NewContainerMap(t *testing.T) {
	cm := NewContainerMap()
	assert.NotNil(t, cm)
	assert.IsType(t, &containerMap{}, cm)
}

func TestContainerMap_Insert(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := mocks.NewMockComponent(ctrl)
	cm := NewContainerMap()

	err := cm.Insert(c)
	assert.NoError(t, err)

	l, err := cm.Len()
	assert.NoError(t, err)
	assert.Equal(t, l, 1)

	// insert nil
	err = cm.Insert(nil)
	assert.Error(t, err)
}

func TestContainerMap_Remove(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := mocks.NewMockComponent(ctrl)
	cm := NewContainerMap()

	cm.Insert(c)
	err := cm.Remove(c)
	assert.NoError(t, err)

	l, err := cm.Len()
	assert.NoError(t, err)
	assert.Equal(t, l, 0)

	// remove thing not in the map
	err = cm.Remove(c)
	assert.NoError(t, err)

	l, err = cm.Len()
	assert.NoError(t, err)
	assert.Equal(t, l, 0)
}

func TestContainerMap_Len(t *testing.T) {
	ctrl := gomock.NewController(t)
	c1 := mocks.NewMockComponent(ctrl)
	c2 := mocks.NewMockComponent(ctrl)
	cm := NewContainerMap()

	l, err := cm.Len()
	assert.NoError(t, err)
	assert.Equal(t, l, 0)

	cm.Insert(c1)
	l, err = cm.Len()
	assert.NoError(t, err)
	assert.Equal(t, l, 1)

	cm.Insert(c2)
	l, err = cm.Len()
	assert.NoError(t, err)
	assert.Equal(t, l, 2)

	// insert things that are already in the map
	cm.Insert(c1)
	cm.Insert(c2)
	l, err = cm.Len()
	assert.NoError(t, err)
	assert.Equal(t, l, 2)

	// remove
	cm.Remove(c1)
	l, err = cm.Len()
	assert.NoError(t, err)
	assert.Equal(t, l, 1)

	cm.Remove(c2)
	l, err = cm.Len()
	assert.NoError(t, err)
	assert.Equal(t, l, 0)
}

func TestContainerMap_Search(t *testing.T) {
	ctrl := gomock.NewController(t)
	c1 := mocks.NewMockComponent(ctrl)
	c2 := mocks.NewMockComponent(ctrl)
	cm := NewContainerMap()

	// search in empty map
	p := utils.Point{}
	c, err := cm.Search(p)
	assert.NoError(t, err)
	assert.Nil(t, c)

	// search c1 in {c1, c2}
	cm.Insert(c1)
	cm.Insert(c2)
	c1.EXPECT().Cover(p).Return(true, nil).Times(1)
	c2.EXPECT().Cover(p).Return(false, nil).Times(1)
	c, err = cm.Search(p)
	assert.NoError(t, err)
	assert.Equal(t, c, c1)

	// search c2 in {c1, c2}
	c1.EXPECT().Cover(p).Return(false, nil).Times(1)
	c2.EXPECT().Cover(p).Return(true, nil).Times(1)
	c, err = cm.Search(p)
	assert.NoError(t, err)
	assert.Equal(t, c, c2)

	// search nil in {c1, c2}
	c1.EXPECT().Cover(p).Return(false, nil).Times(1)
	c2.EXPECT().Cover(p).Return(false, nil).Times(1)
	c, err = cm.Search(p)
	assert.NoError(t, err)
	assert.Nil(t, c)

	// search with layer check
	c1.EXPECT().Cover(p).Return(true, nil).Times(2)
	c2.EXPECT().Cover(p).Return(true, nil).Times(2)
	c1.EXPECT().GetLayer().Return(2, nil).Times(1)
	c2.EXPECT().GetLayer().Return(1, nil).Times(1)
	c, err = cm.Search(p)
	assert.NoError(t, err)
	assert.Equal(t, c, c1)

	c1.EXPECT().GetLayer().Return(1, nil).Times(1)
	c2.EXPECT().GetLayer().Return(2, nil).Times(1)
	c, err = cm.Search(p)
	assert.NoError(t, err)
	assert.Equal(t, c, c2)

	// error in cover
	cm.Remove(c2)
	c1.EXPECT().Cover(p).Return(false, duerror.NewInvalidArgumentError("")).Times(1)
	c, err = cm.Search(p)
	assert.Error(t, err)
	assert.Nil(t, c)

	// error in layer
	cm.Insert(c2)
	c1.EXPECT().Cover(p).Return(true, nil).MaxTimes(1)
	c2.EXPECT().Cover(p).Return(true, nil).MaxTimes(1)
	c1.EXPECT().GetLayer().Return(0, nil).MaxTimes(1)
	c2.EXPECT().GetLayer().Return(0, duerror.NewInvalidArgumentError("")).MaxTimes(1)
	c, err = cm.Search(p)
	assert.Error(t, err)
	assert.Nil(t, c)
}

func TestContainerMap_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	c1 := mocks.NewMockComponent(ctrl)
	c2 := mocks.NewMockComponent(ctrl)
	cm := NewContainerMap()

	// get all in empty map
	cs := cm.GetAll()
	assert.Equal(t, len(cs), 0)

	// get all in {c1, c2}
	cm.Insert(c1)
	cm.Insert(c2)
	cs = cm.GetAll()
	assert.Equal(t, len(cs), 2)

	// remove c1
	cm.Remove(c1)
	cs = cm.GetAll()
	assert.Equal(t, len(cs), 1)

	// remove c2
	cm.Remove(c2)
	cs = cm.GetAll()
	assert.Equal(t, len(cs), 0)
}
