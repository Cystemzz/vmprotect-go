package vmprotect

/*
#include <stdbool.h>
#include "VMProtectSDK.h"
#cgo amd64 LDFLAGS: -L${SRCDIR}/bin -lVMProtectSDK64
#cgo 386 LDFLAGS: -L${SRCDIR}/bin -lVMProtectSDK32
*/
import "C"

func IsProtected() bool {
	return bool(C.VMProtectIsProtected())
}

func IsDebuggerPresent(kernel bool) bool {
	return bool(C.VMProtectIsDebuggerPresent(C.bool(kernel)))
}

func IsVirtualMachinePresent() bool {
	return bool(C.VMProtectIsVirtualMachinePresent())
}

func IsValidImageCRC() bool {
	return bool(C.VMProtectIsValidImageCRC())
}
