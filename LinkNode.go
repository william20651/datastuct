package LinkQueue

type Qnode struct {
	Data interface{}
	Next *Qnode
}

type LinkQueue struct {
	Length int
	Front  *Qnode
	Rear   *Qnode
}

func (lq *LinkQueue) InitSqQueue() *LinkQueue {
	node := &Qnode{}

	return &LinkQueue{
		Length: 0,
		Front:  node,
		Rear:   node,
	}

}

func (lq *LinkQueue) enQueue(data interface{}) *LinkQueue {
	e := &Qnode{
		Data: data,
		Next: nil,
	}
	if lq.Length == 0 {
		lq.Front = e
		lq.Rear = e
	} else {
		lq.Rear.Next = e
		lq.Rear = e
	}
	lq.Length++
	return lq
}

func (lq *LinkQueue) Dequeue() interface{} {
	if lq.Length == 0 {
		return nil
	}

	e := lq.Front
	data := e.Data

	if lq.Length == 1 {
		node := &Qnode{}
		lq.Front = node
		lq.Rear = node
	} else {
		lq.Front = lq.Front.Next
	}
	lq.Length--
	e = nil
	return data
}

func (lq *LinkQueue) LinkQueue2Slice() []interface{} {
	slice := []interface{}{}
	if lq.Length == 0 {
		return nil
	}
	cur := lq.Front
	for cur != nil {
		slice = append(slice, cur.Data)
		if cur.Next != nil {
			cur = cur.Next
		} else {
			break
		}

	}
	return slice
}
