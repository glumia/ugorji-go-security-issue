## ugorji-go-security-issue
Minimal program to reproduce the issue described in fxamacker/cbor's readme in
the *CBOR Library Security* section.

Resources:
- [fxamacker/cbor/README.md#cbor-library-security](https://github.com/fxamacker/cbor/tree/v2.3.0#cbor-security)
- [docs: share test for ugorji/go failing (related discussion on fxamacker/cbor)](https://github.com/fxamacker/cbor/issues/247)
- [oasisprotocol/oasis-core (source of test cases)](https://github.com/oasisprotocol/oasis-core/blob/01ce036d695a5b98075cc90f8063c93802b18f9c/go/common/cbor/cbor_test.go#L10-L24)


#### Execution with `ugorji/go/codec@v1.2.6`

```
gius@graywolf:~/ugorji-go-security-issue$ go test ./ -run TestOutOfMem1
fatal error: runtime: out of memory

runtime stack:
runtime.throw({0x698e0e, 0x303030400000})
        /home/gius/.go/src/runtime/panic.go:1198 +0x71
runtime.sysMap(0xc000400000, 0x42ac40, 0x7ffe76ef2118)
        /home/gius/.go/src/runtime/mem_linux.go:169 +0x96
runtime.(*mheap).grow(0x8b0a20, 0x181818182)
        /home/gius/.go/src/runtime/mheap.go:1393 +0x225
runtime.(*mheap).allocSpan(0x8b0a20, 0x181818182, 0x0, 0x1)
        /home/gius/.go/src/runtime/mheap.go:1179 +0x165
runtime.(*mheap).alloc.func1()
        /home/gius/.go/src/runtime/mheap.go:913 +0x69
runtime.systemstack()
        /home/gius/.go/src/runtime/asm_amd64.s:383 +0x49

goroutine 18 [running]:
runtime.systemstack_switch()
        /home/gius/.go/src/runtime/asm_amd64.s:350 fp=0xc0000423c8 sp=0xc0000423c0 pc=0x466a20
runtime.(*mheap).alloc(0xc000042480, 0x415e0f, 0x0, 0x0)
        /home/gius/.go/src/runtime/mheap.go:907 +0x73 fp=0xc000042418 sp=0xc0000423c8 pc=0x426f73
runtime.(*mcache).allocLarge(0x7f7f5cb0a600, 0x303030303030, 0xab, 0x1)
        /home/gius/.go/src/runtime/mcache.go:227 +0x89 fp=0xc000042478 sp=0xc000042418 pc=0x417be9
runtime.mallocgc(0x303030303030, 0x6470e0, 0x1)
        /home/gius/.go/src/runtime/malloc.go:1082 +0x5c5 fp=0xc0000424f8 sp=0xc000042478 pc=0x40e5a5
runtime.makeslice(0xc0000cc000, 0x203000, 0x24)
        /home/gius/.go/src/runtime/slice.go:98 +0x52 fp=0xc000042520 sp=0xc0000424f8 pc=0x44d952
github.com/ugorji/go/codec.usableByteSlice(...)
        /home/gius/go/pkg/mod/github.com/ugorji/go/codec@v1.1.7/helper.go:2168
github.com/ugorji/go/codec.(*cborDecDriver).DecodeBytes(0xc0000cc000, {0x0, 0x0, 0x0}, 0x0)
        /home/gius/go/pkg/mod/github.com/ugorji/go/codec@v1.1.7/cbor.go:619 +0x145 fp=0xc0000425c0 sp=0xc000042520 pc=0x5bdb45
github.com/ugorji/go/codec.(*Decoder).decode(0xc0000cc040, {0x644300, 0x8c8f40})
        /home/gius/go/pkg/mod/github.com/ugorji/go/codec@v1.1.7/decode.go:1557 +0x2b0 fp=0xc000042670 sp=0xc0000425c0 pc=0x5c5cd0
github.com/ugorji/go/codec.(*Decoder).mustDecode(0xc0000cc040, {0x644300, 0x8c8f40})
        /home/gius/go/pkg/mod/github.com/ugorji/go/codec@v1.1.7/decode.go:1383 +0x38 fp=0xc000042698 sp=0xc000042670 pc=0x5c56d8
github.com/ugorji/go/codec.(*Decoder).Decode(0xc00009e130, {0x644300, 0x8c8f40})
        /home/gius/go/pkg/mod/github.com/ugorji/go/codec@v1.1.7/decode.go:1364 +0xa6 fp=0xc0000426f8 sp=0xc000042698 pc=0x5c5546
github.com/glumia/ugorji-go-security-issue.Unmarshal({0xc00009e130, 0x0, 0x0}, {0x644300, 0x8c8f40})
        /home/gius/ugorji-go-security-issue/cbor.go:9 +0x48 fp=0xc000042730 sp=0xc0000426f8 pc=0x6294c8
github.com/glumia/ugorji-go-security-issue.TestOutOfMem1(0x0)
        /home/gius/ugorji-go-security-issue/cbor_test.go:9 +0x65 fp=0xc000042770 sp=0xc000042730 pc=0x6295a5
testing.tRunner(0xc000082680, 0x6a94c8)
        /home/gius/.go/src/testing/testing.go:1259 +0x102 fp=0xc0000427c0 sp=0xc000042770 pc=0x4dd2c2
testing.(*T).Run·dwrap·21()
        /home/gius/.go/src/testing/testing.go:1306 +0x2a fp=0xc0000427e0 sp=0xc0000427c0 pc=0x4ddfca
runtime.goexit()
        /home/gius/.go/src/runtime/asm_amd64.s:1581 +0x1 fp=0xc0000427e8 sp=0xc0000427e0 pc=0x468c41
created by testing.(*T).Run
        /home/gius/.go/src/testing/testing.go:1306 +0x35a

goroutine 1 [chan receive]:
testing.(*T).Run(0xc0000824e0, {0x695994, 0x46b3d3}, 0x6a94c8)
        /home/gius/.go/src/testing/testing.go:1307 +0x375
testing.runTests.func1(0xc0000824e0)
        /home/gius/.go/src/testing/testing.go:1598 +0x6e
testing.tRunner(0xc0000824e0, 0xc000163d18)
        /home/gius/.go/src/testing/testing.go:1259 +0x102
testing.runTests(0xc0000a0100, {0x889840, 0x2, 0x2}, {0x47f12d, 0x695ffe, 0x8972e0})
        /home/gius/.go/src/testing/testing.go:1596 +0x43f
testing.(*M).Run(0xc0000a0100)
        /home/gius/.go/src/testing/testing.go:1504 +0x51d
main.main()
        _testmain.go:45 +0x14b
FAIL    github.com/glumia/ugorji-go-security-issue      0.196s
FAIL
```

```
gius@graywolf:~/ugorji-go-security-issue$ go test ./ -run TestOutOfMem2
runtime: out of memory: cannot allocate 142048279658496-byte block (3833856 in use)
fatal error: out of memory

goroutine 20 [running]:
runtime.throw({0x695b4e, 0x415e0f})
        /home/gius/.go/src/runtime/panic.go:1198 +0x71 fp=0xc00003c418 sp=0xc00003c3e8 pc=0x436911
runtime.(*mcache).allocLarge(0x7ffa37a62100, 0x813131323233, 0xab, 0x1)
        /home/gius/.go/src/runtime/mcache.go:229 +0x22e fp=0xc00003c478 sp=0xc00003c418 pc=0x417d8e
runtime.mallocgc(0x813131323233, 0x6470e0, 0x1)
        /home/gius/.go/src/runtime/malloc.go:1082 +0x5c5 fp=0xc00003c4f8 sp=0xc00003c478 pc=0x40e5a5
runtime.makeslice(0xc0001a2000, 0x203000, 0x24)
        /home/gius/.go/src/runtime/slice.go:98 +0x52 fp=0xc00003c520 sp=0xc00003c4f8 pc=0x44d952
github.com/ugorji/go/codec.usableByteSlice(...)
        /home/gius/go/pkg/mod/github.com/ugorji/go/codec@v1.1.7/helper.go:2168
github.com/ugorji/go/codec.(*cborDecDriver).DecodeBytes(0xc0001a2000, {0x0, 0x0, 0x0}, 0x0)
        /home/gius/go/pkg/mod/github.com/ugorji/go/codec@v1.1.7/cbor.go:619 +0x145 fp=0xc00003c5c0 sp=0xc00003c520 pc=0x5bdb45
github.com/ugorji/go/codec.(*Decoder).decode(0xc0001a2040, {0x644300, 0x8c8f40})
        /home/gius/go/pkg/mod/github.com/ugorji/go/codec@v1.1.7/decode.go:1557 +0x2b0 fp=0xc00003c670 sp=0xc00003c5c0 pc=0x5c5cd0
github.com/ugorji/go/codec.(*Decoder).mustDecode(0xc0001a2040, {0x644300, 0x8c8f40})
        /home/gius/go/pkg/mod/github.com/ugorji/go/codec@v1.1.7/decode.go:1383 +0x38 fp=0xc00003c698 sp=0xc00003c670 pc=0x5c56d8
github.com/ugorji/go/codec.(*Decoder).Decode(0xc000138480, {0x644300, 0x8c8f40})
        /home/gius/go/pkg/mod/github.com/ugorji/go/codec@v1.1.7/decode.go:1364 +0xa6 fp=0xc00003c6f8 sp=0xc00003c698 pc=0x5c5546
github.com/glumia/ugorji-go-security-issue.Unmarshal({0xc000138480, 0x0, 0x0}, {0x644300, 0x8c8f40})
        /home/gius/ugorji-go-security-issue/cbor.go:9 +0x48 fp=0xc00003c730 sp=0xc00003c6f8 pc=0x6294c8
github.com/glumia/ugorji-go-security-issue.TestOutOfMem2(0x0)
        /home/gius/ugorji-go-security-issue/cbor_test.go:14 +0x65 fp=0xc00003c770 sp=0xc00003c730 pc=0x629625
testing.tRunner(0xc0001a0000, 0x6a94d0)
        /home/gius/.go/src/testing/testing.go:1259 +0x102 fp=0xc00003c7c0 sp=0xc00003c770 pc=0x4dd2c2
testing.(*T).Run·dwrap·21()
        /home/gius/.go/src/testing/testing.go:1306 +0x2a fp=0xc00003c7e0 sp=0xc00003c7c0 pc=0x4ddfca
runtime.goexit()
        /home/gius/.go/src/runtime/asm_amd64.s:1581 +0x1 fp=0xc00003c7e8 sp=0xc00003c7e0 pc=0x468c41
created by testing.(*T).Run
        /home/gius/.go/src/testing/testing.go:1306 +0x35a

goroutine 1 [chan receive]:
testing.(*T).Run(0xc000107d40, {0x6959a1, 0x46b3d3}, 0x6a94d0)
        /home/gius/.go/src/testing/testing.go:1307 +0x375
testing.runTests.func1(0xc000107d40)
        /home/gius/.go/src/testing/testing.go:1598 +0x6e
testing.tRunner(0xc000107d40, 0xc00018fd18)
        /home/gius/.go/src/testing/testing.go:1259 +0x102
testing.runTests(0xc000186100, {0x889840, 0x2, 0x2}, {0x47f12d, 0x695ffe, 0x8972e0})
        /home/gius/.go/src/testing/testing.go:1596 +0x43f
testing.(*M).Run(0xc000186100)
        /home/gius/.go/src/testing/testing.go:1504 +0x51d
main.main()
        _testmain.go:45 +0x14b
FAIL    github.com/glumia/ugorji-go-security-issue      0.005s
FAIL
```


