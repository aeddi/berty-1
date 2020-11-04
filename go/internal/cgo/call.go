package cgo

/*
#cgo LDFLAGS: -lssl -lcrypto -levent

//Encodes Base64
#include <openssl/bio.h>
#include <openssl/evp.h>
#include <openssl/buffer.h>
#include <event2/event.h>
#include <stdint.h>

BUF_MEM *bufferPtr;

const char *base64Encode(const char* text, size_t length) {
	BIO *bio, *b64;

	b64 = BIO_new(BIO_f_base64());
	bio = BIO_new(BIO_s_mem());
	bio = BIO_push(b64, bio);

	BIO_set_flags(bio, BIO_FLAGS_BASE64_NO_NL);
	BIO_write(bio, text, length);
	BIO_flush(bio);
	BIO_get_mem_ptr(bio, &bufferPtr);
	BIO_set_close(bio, BIO_NOCLOSE);
	BIO_free_all(bio);

	return (*bufferPtr).data;
}

int dataLength() {
	return (*bufferPtr).length;
}

void freeBufMem() {
	BUF_MEM_free(bufferPtr);
}

const char* testEvent() {
	struct event_config *cfg = event_config_new();

	event_config_free(cfg);

	return event_get_version();
}
*/
import "C"

import (
	"encoding/base64"
	"fmt"
	"unsafe"
)

func Test(text string) string {
	cText, cLen := C.CString(text), C.ulong(len(text))
	cEncoded := C.base64Encode(cText, cLen)
	encoded := C.GoStringN(cEncoded, C.dataLength())
	C.free(unsafe.Pointer(cText))
	C.freeBufMem()

	bytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(fmt.Sprintf("%s (%s, %s)", err.Error(), text, encoded))
	}

	decoded := string(bytes)
	if text != decoded {
		panic(fmt.Sprintf(
			"text/ decoded mismatch (%s, %s, %s)",
			text,
			encoded,
			decoded,
		))
	}

	eventVersion := C.GoString(C.testEvent())

	return fmt.Sprintf(
		"\ttext: %s\n\tencoded: %s\n\tdecoded: %s\n\tlibevent version: %s",
		text,
		encoded,
		decoded,
		eventVersion,
	)
}
