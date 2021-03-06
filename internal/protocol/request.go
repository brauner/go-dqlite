package protocol

// DO NOT EDIT
//
// This file was generated by ./schema.sh


// EncodeLeader encodes a Leader request.
func EncodeLeader(request *Message) {
	request.reset()
	request.putUint64(0)

	request.putHeader(RequestLeader)
}

// EncodeClient encodes a Client request.
func EncodeClient(request *Message, id uint64) {
	request.reset()
	request.putUint64(id)

	request.putHeader(RequestClient)
}

// EncodeHeartbeat encodes a Heartbeat request.
func EncodeHeartbeat(request *Message, timestamp uint64) {
	request.reset()
	request.putUint64(timestamp)

	request.putHeader(RequestHeartbeat)
}

// EncodeOpen encodes a Open request.
func EncodeOpen(request *Message, name string, flags uint64, vfs string) {
	request.reset()
	request.putString(name)
	request.putUint64(flags)
	request.putString(vfs)

	request.putHeader(RequestOpen)
}

// EncodePrepare encodes a Prepare request.
func EncodePrepare(request *Message, db uint64, sql string) {
	request.reset()
	request.putUint64(db)
	request.putString(sql)

	request.putHeader(RequestPrepare)
}

// EncodeExec encodes a Exec request.
func EncodeExec(request *Message, db uint32, stmt uint32, values NamedValues) {
	request.reset()
	request.putUint32(db)
	request.putUint32(stmt)
	request.putNamedValues(values)

	request.putHeader(RequestExec)
}

// EncodeQuery encodes a Query request.
func EncodeQuery(request *Message, db uint32, stmt uint32, values NamedValues) {
	request.reset()
	request.putUint32(db)
	request.putUint32(stmt)
	request.putNamedValues(values)

	request.putHeader(RequestQuery)
}

// EncodeFinalize encodes a Finalize request.
func EncodeFinalize(request *Message, db uint32, stmt uint32) {
	request.reset()
	request.putUint32(db)
	request.putUint32(stmt)

	request.putHeader(RequestFinalize)
}

// EncodeExecSQL encodes a ExecSQL request.
func EncodeExecSQL(request *Message, db uint64, sql string, values NamedValues) {
	request.reset()
	request.putUint64(db)
	request.putString(sql)
	request.putNamedValues(values)

	request.putHeader(RequestExecSQL)
}

// EncodeQuerySQL encodes a QuerySQL request.
func EncodeQuerySQL(request *Message, db uint64, sql string, values NamedValues) {
	request.reset()
	request.putUint64(db)
	request.putString(sql)
	request.putNamedValues(values)

	request.putHeader(RequestQuerySQL)
}

// EncodeInterrupt encodes a Interrupt request.
func EncodeInterrupt(request *Message, db uint64) {
	request.reset()
	request.putUint64(db)

	request.putHeader(RequestInterrupt)
}

// EncodeAdd encodes a Add request.
func EncodeAdd(request *Message, id uint64, address string) {
	request.reset()
	request.putUint64(id)
	request.putString(address)

	request.putHeader(RequestAdd)
}

// EncodeAssign encodes a Assign request.
func EncodeAssign(request *Message, id uint64, role uint64) {
	request.reset()
	request.putUint64(id)
	request.putUint64(role)

	request.putHeader(RequestAssign)
}

// EncodeRemove encodes a Remove request.
func EncodeRemove(request *Message, id uint64) {
	request.reset()
	request.putUint64(id)

	request.putHeader(RequestRemove)
}

// EncodeDump encodes a Dump request.
func EncodeDump(request *Message, name string) {
	request.reset()
	request.putString(name)

	request.putHeader(RequestDump)
}

// EncodeCluster encodes a Cluster request.
func EncodeCluster(request *Message, format uint64) {
	request.reset()
	request.putUint64(format)

	request.putHeader(RequestCluster)
}

// EncodeTransfer encodes a Transfer request.
func EncodeTransfer(request *Message, id uint64) {
	request.reset()
	request.putUint64(id)

	request.putHeader(RequestTransfer)
}

// EncodeDescribe encodes a Describe request.
func EncodeDescribe(request *Message, format uint64) {
	request.reset()
	request.putUint64(format)

	request.putHeader(RequestDescribe)
}

// EncodeWeight encodes a Weight request.
func EncodeWeight(request *Message, weight uint64) {
	request.reset()
	request.putUint64(weight)

	request.putHeader(RequestWeight)
}
