// Code generated by 'ccgo sys/stat/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o sys/stat/stat_windows_386.go -pkgname stat', DO NOT EDIT.

package stat

import (
	"math"
	"reflect"
	"sync/atomic"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ atomic.Value
var _ unsafe.Pointer

const (
	DUMMYSTRUCTNAME                                 = 0          // _mingw.h:519:1:
	DUMMYSTRUCTNAME1                                = 0          // _mingw.h:520:1:
	DUMMYSTRUCTNAME2                                = 0          // _mingw.h:521:1:
	DUMMYSTRUCTNAME3                                = 0          // _mingw.h:522:1:
	DUMMYSTRUCTNAME4                                = 0          // _mingw.h:523:1:
	DUMMYSTRUCTNAME5                                = 0          // _mingw.h:524:1:
	DUMMYUNIONNAME                                  = 0          // _mingw.h:497:1:
	DUMMYUNIONNAME1                                 = 0          // _mingw.h:498:1:
	DUMMYUNIONNAME2                                 = 0          // _mingw.h:499:1:
	DUMMYUNIONNAME3                                 = 0          // _mingw.h:500:1:
	DUMMYUNIONNAME4                                 = 0          // _mingw.h:501:1:
	DUMMYUNIONNAME5                                 = 0          // _mingw.h:502:1:
	DUMMYUNIONNAME6                                 = 0          // _mingw.h:503:1:
	DUMMYUNIONNAME7                                 = 0          // _mingw.h:504:1:
	DUMMYUNIONNAME8                                 = 0          // _mingw.h:505:1:
	DUMMYUNIONNAME9                                 = 0          // _mingw.h:506:1:
	F_OK                                            = 0          // io.h:182:1:
	MINGW_DDK_H                                     = 0          // _mingw_ddk.h:2:1:
	MINGW_HAS_DDK_H                                 = 1          // _mingw_ddk.h:4:1:
	MINGW_HAS_SECURE_API                            = 1          // _mingw.h:602:1:
	MINGW_SDK_INIT                                  = 0          // _mingw.h:598:1:
	R_OK                                            = 4          // io.h:185:1:
	S_IEXEC                                         = 64         // stat.h:141:1:
	S_IFBLK                                         = 12288      // stat.h:143:1:
	S_IFCHR                                         = 8192       // stat.h:137:1:
	S_IFDIR                                         = 16384      // stat.h:136:1:
	S_IFIFO                                         = 4096       // stat.h:142:1:
	S_IFMT                                          = 61440      // stat.h:135:1:
	S_IFREG                                         = 32768      // stat.h:138:1:
	S_IREAD                                         = 256        // stat.h:139:1:
	S_IRGRP                                         = 32         // stat.h:155:1:
	S_IROTH                                         = 4          // stat.h:160:1:
	S_IRUSR                                         = 256        // stat.h:152:1:
	S_IRWXG                                         = 56         // stat.h:158:1:
	S_IRWXO                                         = 7          // stat.h:163:1:
	S_IRWXU                                         = 448        // stat.h:149:1:
	S_IWGRP                                         = 16         // stat.h:156:1:
	S_IWOTH                                         = 2          // stat.h:161:1:
	S_IWRITE                                        = 128        // stat.h:140:1:
	S_IWUSR                                         = 128        // stat.h:151:1:
	S_IXGRP                                         = 8          // stat.h:157:1:
	S_IXOTH                                         = 1          // stat.h:162:1:
	S_IXUSR                                         = 64         // stat.h:150:1:
	UNALIGNED                                       = 0          // _mingw.h:384:1:
	USE___UUIDOF                                    = 0          // _mingw.h:77:1:
	WIN32                                           = 1          // <predefined>:258:1:
	WINNT                                           = 1          // <predefined>:306:1:
	W_OK                                            = 2          // io.h:184:1:
	X_OK                                            = 1          // io.h:183:1:
	X_AGLOBAL                                       = 0          // _mingw.h:346:1:
	X_ANONYMOUS_STRUCT                              = 0          // _mingw.h:474:1:
	X_ANONYMOUS_UNION                               = 0          // _mingw.h:473:1:
	X_ARGMAX                                        = 100        // _mingw.h:402:1:
	X_A_ARCH                                        = 0x20       // io.h:156:1:
	X_A_HIDDEN                                      = 0x02       // io.h:153:1:
	X_A_NORMAL                                      = 0x00       // io.h:151:1:
	X_A_RDONLY                                      = 0x01       // io.h:152:1:
	X_A_SUBDIR                                      = 0x10       // io.h:155:1:
	X_A_SYSTEM                                      = 0x04       // io.h:154:1:
	X_CONST_RETURN                                  = 0          // _mingw.h:377:1:
	X_CRTNOALIAS                                    = 0          // corecrt.h:29:1:
	X_CRTRESTRICT                                   = 0          // corecrt.h:33:1:
	X_CRT_ALTERNATIVE_IMPORTED                      = 0          // _mingw.h:313:1:
	X_CRT_DIRECTORY_DEFINED                         = 0          // io.h:214:1:
	X_CRT_MANAGED_HEAP_DEPRECATE                    = 0          // _mingw.h:361:1:
	X_CRT_MEMORY_DEFINED                            = 0          // string.h:44:1:
	X_CRT_PACKING                                   = 8          // corecrt.h:14:1:
	X_CRT_SECURE_CPP_OVERLOAD_SECURE_NAMES          = 0          // _mingw_secapi.h:34:1:
	X_CRT_SECURE_CPP_OVERLOAD_SECURE_NAMES_MEMORY   = 0          // _mingw_secapi.h:35:1:
	X_CRT_SECURE_CPP_OVERLOAD_STANDARD_NAMES        = 0          // _mingw_secapi.h:36:1:
	X_CRT_SECURE_CPP_OVERLOAD_STANDARD_NAMES_COUNT  = 0          // _mingw_secapi.h:37:1:
	X_CRT_SECURE_CPP_OVERLOAD_STANDARD_NAMES_MEMORY = 0          // _mingw_secapi.h:38:1:
	X_CRT_USE_WINAPI_FAMILY_DESKTOP_APP             = 0          // corecrt.h:501:1:
	X_DEV_T_DEFINED                                 = 0          // types.h:50:1:
	X_DLL                                           = 0          // _mingw.h:326:1:
	X_ERRCODE_DEFINED                               = 0          // corecrt.h:117:1:
	X_FILE_OFFSET_BITS                              = 64         // <builtin>:25:1:
	X_FILE_OFFSET_BITS_SET_LSEEK                    = 0          // io.h:350:1:
	X_FILE_OFFSET_BITS_SET_OFFT                     = 0          // _mingw_off_t.h:21:1:
	X_FINDDATA_T_DEFINED                            = 0          // io.h:89:1:
	X_FSIZE_T_DEFINED                               = 0          // io.h:30:1:
	X_ILP32                                         = 1          // <predefined>:211:1:
	X_INC_CORECRT                                   = 0          // corecrt.h:8:1:
	X_INC_CRTDEFS                                   = 0          // crtdefs.h:8:1:
	X_INC_CRTDEFS_MACRO                             = 0          // _mingw_mac.h:8:1:
	X_INC_MINGW_SECAPI                              = 0          // _mingw_secapi.h:8:1:
	X_INC_STAT                                      = 0          // stat.h:7:1:
	X_INC_STRING                                    = 0          // string.h:7:1:
	X_INC_STRING_S                                  = 0          // string_s.h:7:1:
	X_INC_TYPES                                     = 0          // types.h:7:1:
	X_INC_VADEFS                                    = 0          // vadefs.h:7:1:
	X_INC__MINGW_H                                  = 0          // _mingw.h:8:1:
	X_INO_T_DEFINED                                 = 0          // types.h:42:1:
	X_INT128_DEFINED                                = 0          // _mingw.h:237:1:
	X_INTEGRAL_MAX_BITS                             = 64         // <predefined>:320:1:
	X_INTPTR_T_DEFINED                              = 0          // corecrt.h:62:1:
	X_IO_H_                                         = 0          // io.h:7:1:
	X_MODE_T_                                       = 0          // types.h:73:1:
	X_MT                                            = 0          // _mingw.h:330:1:
	X_M_IX86                                        = 600        // _mingw_mac.h:54:1:
	X_NLSCMPERROR                                   = 2147483647 // string.h:26:1:
	X_NLSCMP_DEFINED                                = 0          // string.h:25:1:
	X_OFF64_T_DEFINED                               = 0          // _mingw_off_t.h:12:1:
	X_OFF_T_                                        = 0          // _mingw_off_t.h:4:1:
	X_OFF_T_DEFINED                                 = 0          // _mingw_off_t.h:2:1:
	X_PGLOBAL                                       = 0          // _mingw.h:342:1:
	X_PID_T_                                        = 0          // types.h:58:1:
	X_PTRDIFF_T_                                    = 0          // corecrt.h:90:1:
	X_PTRDIFF_T_DEFINED                             = 0          // corecrt.h:88:1:
	X_RSIZE_T_DEFINED                               = 0          // corecrt.h:58:1:
	X_SECURECRT_FILL_BUFFER_PATTERN                 = 0xFD       // _mingw.h:349:1:
	X_SIGSET_T_                                     = 0          // types.h:101:1:
	X_SIZE_T_DEFINED                                = 0          // corecrt.h:37:1:
	X_SSIZE_T_DEFINED                               = 0          // corecrt.h:47:1:
	X_STAT_DEFINED                                  = 0          // _mingw_stat64.h:101:1:
	X_S_IEXEC                                       = 0x0040     // stat.h:67:1:
	X_S_IFBLK                                       = 0x3000     // stat.h:133:1:
	X_S_IFCHR                                       = 0x2000     // stat.h:62:1:
	X_S_IFDIR                                       = 0x4000     // stat.h:61:1:
	X_S_IFIFO                                       = 0x1000     // stat.h:63:1:
	X_S_IFMT                                        = 0xF000     // stat.h:60:1:
	X_S_IFREG                                       = 0x8000     // stat.h:64:1:
	X_S_IREAD                                       = 0x0100     // stat.h:65:1:
	X_S_IRUSR                                       = 256        // stat.h:153:1:
	X_S_IRWXU                                       = 448        // stat.h:145:1:
	X_S_IWRITE                                      = 0x0080     // stat.h:66:1:
	X_S_IWUSR                                       = 128        // stat.h:147:1:
	X_S_IXUSR                                       = 64         // stat.h:146:1:
	X_TAGLC_ID_DEFINED                              = 0          // corecrt.h:447:1:
	X_THREADLOCALEINFO                              = 0          // corecrt.h:456:1:
	X_TIME32_T_DEFINED                              = 0          // corecrt.h:122:1:
	X_TIME64_T_DEFINED                              = 0          // corecrt.h:127:1:
	X_TIMESPEC_DEFINED                              = 0          // types.h:88:1:
	X_TIME_T_DEFINED                                = 0          // corecrt.h:139:1:
	X_UINTPTR_T_DEFINED                             = 0          // corecrt.h:75:1:
	X_USE_32BIT_TIME_T                              = 0          // _mingw.h:372:1:
	X_VA_LIST_DEFINED                               = 0          // <builtin>:55:1:
	X_W64                                           = 0          // _mingw.h:296:1:
	X_WCHAR_T_DEFINED                               = 0          // corecrt.h:101:1:
	X_WCTYPE_T_DEFINED                              = 0          // corecrt.h:108:1:
	X_WConst_return                                 = 0          // string.h:41:1:
	X_WFINDDATA_T_DEFINED                           = 0          // io.h:148:1:
	X_WIN32                                         = 1          // <predefined>:164:1:
	X_WIN32_WINNT                                   = 0x502      // _mingw.h:233:1:
	X_WINT_T                                        = 0          // corecrt.h:110:1:
	X_WIO_DEFINED                                   = 0          // io.h:295:1:
	X_WSTAT_DEFINED                                 = 0          // stat.h:125:1:
	X_WSTRING_DEFINED                               = 0          // string.h:129:1:
	X_WSTRING_S_DEFINED                             = 0          // string_s.h:48:1:
	X_X86_                                          = 1          // <predefined>:169:1:
	I386                                            = 1          // <predefined>:171:1:
)

type Ptrdiff_t = int32 /* <builtin>:3:26 */

type Size_t = uint32 /* <builtin>:9:23 */

type Wchar_t = uint16 /* <builtin>:15:24 */

type X__builtin_va_list = uintptr /* <builtin>:46:14 */
type X__float128 = float64        /* <builtin>:47:21 */

type Va_list = X__builtin_va_list /* <builtin>:50:27 */

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// This macro holds an monotonic increasing value, which indicates
//    a specific fix/patch is present on trunk.  This value isn't related to
//    minor/major version-macros.  It is increased on demand, if a big
//    fix was applied to trunk.  This macro gets just increased on trunk.  For
//    other branches its value won't be modified.

// mingw.org's version macros: these make gcc to define
//    MINGW32_SUPPORTS_MT_EH and to use the _CRT_MT global
//    and the __mingwthr_key_dtor() function from the MinGW
//    CRT in its private gthr-win32.h header.

// Set VC specific compiler target macros.

// For x86 we have always to prefix by underscore.

// Special case nameless struct/union.

// MinGW-w64 has some additional C99 printf/scanf feature support.
//    So we add some helper macros to ease recognition of them.

// If _FORTIFY_SOURCE is enabled, some inline functions may use
//    __builtin_va_arg_pack().  GCC may report an error if the address
//    of such a function is used.  Set _FORTIFY_VA_ARG=0 in this case.

// Enable workaround for ABI incompatibility on affected platforms

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// http://msdn.microsoft.com/en-us/library/ms175759%28v=VS.100%29.aspx
// Templates won't work in C, will break if secure API is not enabled, disabled

// https://blogs.msdn.com/b/sdl/archive/2010/02/16/vc-2010-and-memcpy.aspx?Redirected=true
// fallback on default implementation if we can't know the size of the destination

// Include _cygwin.h if we're building a Cygwin application.

// Target specific macro replacement for type "long".  In the Windows API,
//    the type long is always 32 bit, even if the target is 64 bit (LLP64).
//    On 64 bit Cygwin, the type long is 64 bit (LP64).  So, to get the right
//    sized definitions and declarations, all usage of type long in the Windows
//    headers have to be replaced by the below defined macro __LONG32.

// C/C++ specific language defines.

// Note the extern. This is needed to work around GCC's
// limitations in handling dllimport attribute.

// Attribute `nonnull' was valid as of gcc 3.3.  We don't use GCC's
//    variadiac macro facility, because variadic macros cause syntax
//    errors with  --traditional-cpp.

//  High byte is the major version, low byte is the minor.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// for backward compatibility

type X__gnuc_va_list = X__builtin_va_list /* vadefs.h:24:29 */

type Ssize_t = int32 /* corecrt.h:52:13 */

type Rsize_t = Size_t /* corecrt.h:57:16 */

type Intptr_t = int32 /* corecrt.h:69:13 */

type Uintptr_t = uint32 /* corecrt.h:82:22 */

type Wint_t = uint16   /* corecrt.h:111:24 */
type Wctype_t = uint16 /* corecrt.h:112:24 */

type Errno_t = int32 /* corecrt.h:118:13 */

type X__time32_t = int32 /* corecrt.h:123:14 */

type X__time64_t = int64 /* corecrt.h:128:35 */

type Time_t = X__time32_t /* corecrt.h:141:20 */

type Threadlocaleinfostruct = struct {
	Frefcount      int32
	Flc_codepage   uint32
	Flc_collate_cp uint32
	Flc_handle     [6]uint32
	Flc_id         [6]LC_ID
	Flc_category   [6]struct {
		Flocale    uintptr
		Fwlocale   uintptr
		Frefcount  uintptr
		Fwrefcount uintptr
	}
	Flc_clike            int32
	Fmb_cur_max          int32
	Flconv_intl_refcount uintptr
	Flconv_num_refcount  uintptr
	Flconv_mon_refcount  uintptr
	Flconv               uintptr
	Fctype1_refcount     uintptr
	Fctype1              uintptr
	Fpctype              uintptr
	Fpclmap              uintptr
	Fpcumap              uintptr
	Flc_time_curr        uintptr
} /* corecrt.h:435:1 */

type Pthreadlocinfo = uintptr /* corecrt.h:437:39 */
type Pthreadmbcinfo = uintptr /* corecrt.h:438:36 */

type Localeinfo_struct = struct {
	Flocinfo Pthreadlocinfo
	Fmbcinfo Pthreadmbcinfo
} /* corecrt.h:441:9 */

type X_locale_tstruct = Localeinfo_struct /* corecrt.h:444:3 */
type X_locale_t = uintptr                 /* corecrt.h:444:19 */

type TagLC_ID = struct {
	FwLanguage uint16
	FwCountry  uint16
	FwCodePage uint16
} /* corecrt.h:435:1 */

type LC_ID = TagLC_ID  /* corecrt.h:452:3 */
type LPLC_ID = uintptr /* corecrt.h:452:9 */

type Threadlocinfo = Threadlocaleinfostruct /* corecrt.h:487:3 */
type X_fsize_t = uint32                     /* io.h:29:25 */

type X_finddata32_t = struct {
	Fattrib      uint32
	Ftime_create X__time32_t
	Ftime_access X__time32_t
	Ftime_write  X__time32_t
	Fsize        X_fsize_t
	Fname        [260]int8
} /* io.h:35:3 */

type X_finddata32i64_t = struct {
	Fattrib      uint32
	Ftime_create X__time32_t
	Ftime_access X__time32_t
	Ftime_write  X__time32_t
	Fsize        int64
	Fname        [260]int8
	F__ccgo_pad1 [4]byte
} /* io.h:44:3 */

type X_finddata64i32_t = struct {
	Fattrib      uint32
	F__ccgo_pad1 [4]byte
	Ftime_create X__time64_t
	Ftime_access X__time64_t
	Ftime_write  X__time64_t
	Fsize        X_fsize_t
	Fname        [260]int8
} /* io.h:53:3 */

type X__finddata64_t = struct {
	Fattrib      uint32
	F__ccgo_pad1 [4]byte
	Ftime_create X__time64_t
	Ftime_access X__time64_t
	Ftime_write  X__time64_t
	Fsize        int64
	Fname        [260]int8
	F__ccgo_pad2 [4]byte
} /* io.h:62:3 */

type X_wfinddata32_t = struct {
	Fattrib      uint32
	Ftime_create X__time32_t
	Ftime_access X__time32_t
	Ftime_write  X__time32_t
	Fsize        X_fsize_t
	Fname        [260]Wchar_t
} /* io.h:94:3 */

type X_wfinddata32i64_t = struct {
	Fattrib      uint32
	Ftime_create X__time32_t
	Ftime_access X__time32_t
	Ftime_write  X__time32_t
	Fsize        int64
	Fname        [260]Wchar_t
} /* io.h:103:3 */

type X_wfinddata64i32_t = struct {
	Fattrib      uint32
	F__ccgo_pad1 [4]byte
	Ftime_create X__time64_t
	Ftime_access X__time64_t
	Ftime_write  X__time64_t
	Fsize        X_fsize_t
	Fname        [260]Wchar_t
	F__ccgo_pad2 [4]byte
} /* io.h:112:3 */

type X_wfinddata64_t = struct {
	Fattrib      uint32
	F__ccgo_pad1 [4]byte
	Ftime_create X__time64_t
	Ftime_access X__time64_t
	Ftime_write  X__time64_t
	Fsize        int64
	Fname        [260]Wchar_t
} /* io.h:121:3 */

type X_off_t = int32 /* _mingw_off_t.h:5:16 */
type Off32_t = int32 /* _mingw_off_t.h:7:16 */

type X_off64_t = int64 /* _mingw_off_t.h:13:39 */
type Off64_t = int64   /* _mingw_off_t.h:15:39 */

type Off_t = Off64_t /* _mingw_off_t.h:24:17 */

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

type X_ino_t = uint16 /* types.h:43:24 */
type Ino_t = uint16   /* types.h:45:24 */

type X_dev_t = uint32 /* types.h:51:22 */
type Dev_t = uint32   /* types.h:53:22 */

type X_pid_t = int32 /* types.h:60:13 */

type Pid_t = X_pid_t /* types.h:68:16 */

type X_mode_t = uint16 /* types.h:74:24 */

type Mode_t = X_mode_t /* types.h:77:17 */

type Useconds_t = uint32 /* types.h:84:22 */

type Timespec = struct {
	Ftv_sec  Time_t
	Ftv_nsec int32
} /* types.h:89:1 */

type Itimerspec = struct {
	Fit_interval struct {
		Ftv_sec  Time_t
		Ftv_nsec int32
	}
	Fit_value struct {
		Ftv_sec  Time_t
		Ftv_nsec int32
	}
} /* types.h:94:1 */

type X_sigset_t = uint32 /* types.h:106:23 */

type X_stat32 = struct {
	Fst_dev      X_dev_t
	Fst_ino      X_ino_t
	Fst_mode     uint16
	Fst_nlink    int16
	Fst_uid      int16
	Fst_gid      int16
	F__ccgo_pad1 [2]byte
	Fst_rdev     X_dev_t
	Fst_size     X_off_t
	Fst_atime    X__time32_t
	Fst_mtime    X__time32_t
	Fst_ctime    X__time32_t
} /* _mingw_stat64.h:25:3 */

type Stat = struct {
	Fst_dev      X_dev_t
	Fst_ino      X_ino_t
	Fst_mode     uint16
	Fst_nlink    int16
	Fst_uid      int16
	Fst_gid      int16
	F__ccgo_pad1 [2]byte
	Fst_rdev     X_dev_t
	Fst_size     X_off_t
	Fst_atime    Time_t
	Fst_mtime    Time_t
	Fst_ctime    Time_t
} /* _mingw_stat64.h:40:3 */

type X_stati64 = struct {
	Fst_dev      X_dev_t
	Fst_ino      X_ino_t
	Fst_mode     uint16
	Fst_nlink    int16
	Fst_uid      int16
	Fst_gid      int16
	F__ccgo_pad1 [2]byte
	Fst_rdev     X_dev_t
	F__ccgo_pad2 [4]byte
	Fst_size     int64
	Fst_atime    X__time32_t
	Fst_mtime    X__time32_t
	Fst_ctime    X__time32_t
	F__ccgo_pad3 [4]byte
} /* _mingw_stat64.h:55:3 */

type X_stat64i32 = struct {
	Fst_dev      X_dev_t
	Fst_ino      X_ino_t
	Fst_mode     uint16
	Fst_nlink    int16
	Fst_uid      int16
	Fst_gid      int16
	F__ccgo_pad1 [2]byte
	Fst_rdev     X_dev_t
	Fst_size     X_off_t
	Fst_atime    X__time64_t
	Fst_mtime    X__time64_t
	Fst_ctime    X__time64_t
} /* _mingw_stat64.h:69:3 */

type X_stat64 = struct {
	Fst_dev      X_dev_t
	Fst_ino      X_ino_t
	Fst_mode     uint16
	Fst_nlink    int16
	Fst_uid      int16
	Fst_gid      int16
	F__ccgo_pad1 [2]byte
	Fst_rdev     X_dev_t
	F__ccgo_pad2 [4]byte
	Fst_size     int64
	Fst_atime    X__time64_t
	Fst_mtime    X__time64_t
	Fst_ctime    X__time64_t
} /* _mingw_stat64.h:83:3 */

var _ int8 /* gen.c:2:13: */