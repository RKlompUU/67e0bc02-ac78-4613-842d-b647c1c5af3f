This is a simple HTTP server, that serves data from a database that contains some
fraction of Ethereum onchain data.

# Running locally (optional)

Start the server:

docker-compose up --build


You should now be able to call the API:

curl http://localhost:8080/metrics/ethereum/transaction_fees_hourly?date=2020-09-07


## Open issue

API calls return with 500, with the server printing an error message containing:
... runtime error: invalid memory address or nil pointer dereference ... pkg/metrics.go:44 ...

{"time":"2025-03-28T12:36:50.449570129Z","level":"-","prefix":"echo","file""recover.go","line":"92","message":"[PANIC RECOVER] runtime error: invalid memory address or nil pointer derefernce goroutine 8 [running]:\nmain.NewServer.Recover.RecoverWithConfig.func3.1.1()\n\t/go/pkg/mod/github.com/labstck/echo/v4@v4.6.1/middleware/recover.go:77 +0xd8\npanic({0x38d200?, 0x6e7a10?})\n\t/usr/local/go/src/runtime/panc.go:914 +0x218\ngithub.com/jackc/pgx/v4/pgxpool.(*Pool).Acquire(0x0, {0x4949d0, 0x72cfe0})\n\t/go/pkg/mod/githu.com/jackc/pgx/v4@v4.14.1/pgxpool/pool.go:414 +0x38\ngithub.com/jackc/pgx/v4/pgxpool.(*Pool).Query(0x40002256b8? {0x4949d0, 0x72cfe0}, {0x40fb2b, 0x155}, {0x400020e050, 0x1, 0x1})\n\t/go/pkg/mod/github.com/jackc/pgx/v4@v4.141/pgxpool/pool.go:491 +0x3c\ngithub.com/georgysavva/scany/pgxscan.Select({0x4949d0?, 0x72cfe0?}, {0x490440?, 0x400040298?}, {0x371620, 0x4000240000}, {0x40fb2b?, 0x0?}, {0x400020e050?, 0x4000232168?, ...})\n\t/go/pkg/mod/gitub.com/georgysavva/scany@v0.2.9/pgxscan/pgxscan.go:28 +0x60\nmain.(*Database).GetEOATransactionFeesHourly(0x3f5b3?, {0x0, 0xed6e76f00, 0x0})\n\t/svc/pkg/metrics.go:44 +0xd0\nmain.NewServer.TransferFeesHourly.func2({0x4987f0,0x400023e000})\n\t/svc/pkg/handler.go:50 +0x90\ngithub.com/labstack/echo/v4.(*Echo).add.func1({0x4987f0, 0x40002e000})\n\t/go/pkg/mod/github.com/labstack/echo/v4@v4.6.1/echo.go:552 +0x50\nmain.NewServer.Recover.RecoverWithCofig.func3.1({0x4987f0?, 0x400023e000})\n\t/go/pkg/mod/github.com/labstack/echo/v4@v4.6.1/middleware/recover.go:9 +0xd8\ngithub.com/labstack/echo/v4/middleware.LoggerWithConfig.func2.1({0x4987f0?, 0x400023e000})\n\t/go/pkg/mo/github.com/labstack/echo/v4@v4.6.1/middleware/logger.go:117 +0xa4\ngithub.com/labstack/echo/v4.(*Echo).ServeHTT(0x400014efc0, {0x492d60?, 0x4000238000}, 0x400021a000)\n\t/go/pkg/mod/github.com/labstack/echo/v4@v4.6.1/echo.g:662 +0x374\nnet/http.serverHandler.ServeHTTP({0x400020a030?}, {0x492d60?, 0x4000238000?}, 0x6?)\n\t/usr/local/g/src/net/http/server.go:2943 +0xbc\nnet/http.(*conn).serve(0x4000146360, {0x494a08, 0x4000188a80})\n\t/usr/localgo/src/net/http/server.go:2014 +0x518\ncreated by net/http.(*Server).Serve in goroutine 1\n\t/usr/local/go/src/nt/http/server.go:3091 +0x4cc\n\ngoroutine 1 [IO wait]:\ninternal/poll.runtime_pollWait(0xffff4f2501a0, 0x72)\n\tusr/local/go/src/runtime/netpoll.go:343 +0xa0\ninternal/poll.(*pollDesc).wait(0x4000021100?, 0x24290?, 0x0)\n\t/sr/local/go/src/internal/poll/fd_poll_runtime.go:84 +0x28\ninternal/poll.(*pollDesc).waitRead(...)\n\t/usr/localgo/src/internal/poll/fd_poll_runtime.go:89\ninternal/poll.(*FD).Accept(0x4000021100)\n\t/usr/local/go/src/internl/poll/fd_unix.go:611 +0x250\nnet.(*netFD).accept(0x4000021100)\n\t/usr/local/go/src/net/fd_unix.go:172 +0x28\nnt.(*TCPListener).accept(0x400005efc0)\n\t/usr/local/go/src/net/tcpsock_posix.go:152 +0x28\nnet.(*TCPListener).AceptTCP(0x400005efc0)\n\t/usr/local/go/src/net/tcpsock.go:302 +0x2c\ngithub.com/labstack/echo/v4.tcpKeepAliveListner.Accept({0x400017bce8?})\n\t/go/pkg/mod/github.com/labstack/echo/v4@v4.6.1/echo.go:971 +0x1c\nnet/http.(*Servr).Serve(0x400018c1e0, {0x492e80, 0x40000402c8})\n\t/usr/local/go/src/net/http/server.go:3061 +0x2b8\ngithub.comlabstack/echo/v4.(*Echo).Start(0x400014efc0, {0x400000f196, 0x5})\n\t/go/pkg/mod/github.com/labstack/echo/v4@v4..1/echo.go:679 +0xc0\nmain.main()\n\t/svc/pkg/main.go:19 +0x12c\n\ngoroutine 17 [runnable]:\nsyscall.Syscall(0x40003edd8?, 0xc6980?, 0x800000?, 0x7ffff800000?)\n\t/usr/local/go/src/syscall/syscall_linux.go:69 +0x20\nsyscall.ead(0x4000021180?, {0x400020a041?, 0x0?, 0x0?})\n\t/usr/local/go/src/syscall/zsyscall_linux_arm64.go:721 +0x40\nyscall.Read(...)\n\t/usr/local/go/src/syscall/syscall_unix.go:181\ninternal/poll.ignoringEINTRIO(...)\n\t/usr/loal/go/src/internal/poll/fd_unix.go:736\ninternal/poll.(*FD).Read(0x4000021180, {0x400020a041, 0x1, 0x1})\n\t/usrlocal/go/src/internal/poll/fd_unix.go:160 +0x224\nnet.(*netFD).Read(0x4000021180, {0x400020a041?, 0x0?, 0x0?})\nt/usr/local/go/src/net/fd_posix.go:55 +0x28\nnet.(*conn).Read(0x40000402e0, {0x400020a041?, 0x0?, 0x0?})\n\t/usrlocal/go/src/net/net.go:185 +0x34\nnet/http.(*connReader).backgroundRead(0x400020a030)\n\t/usr/local/go/src/net/ttp/server.go:683 +0x40\ncreated by net/http.(*connReader).startBackgroundRead in goroutine 8\n\t/usr/local/go/sc/net/http/server.go:679 \n"}



While reading the codebase it would be nice if you can:
- explain what you're seeing
- how you're interpreting the codebase
- spot issues and explain how you would fix them
