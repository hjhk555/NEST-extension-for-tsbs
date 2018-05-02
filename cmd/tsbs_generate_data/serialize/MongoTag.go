// automatically generated by the FlatBuffers compiler, do not modify

package serialize

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MongoTag struct {
	_tab flatbuffers.Table
}

func GetRootAsMongoTag(buf []byte, offset flatbuffers.UOffsetT) *MongoTag {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MongoTag{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *MongoTag) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MongoTag) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MongoTag) Key() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *MongoTag) Value() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func MongoTagStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func MongoTagAddKey(builder *flatbuffers.Builder, key flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(key), 0)
}
func MongoTagAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(value), 0)
}
func MongoTagEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}