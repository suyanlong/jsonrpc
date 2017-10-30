--
-- Created by IntelliJ IDEA.
-- User: suyanlong
-- Date: 17-10-30
-- Time: 下午5:12
-- To change this template use File | Settings | File Templates.
-- run:
-- wrk -c 100 -d 10 -t 32 -s post.lua http://127.0.0.1:1337

wrk.method = "POST"
wrk.body = '{"jsonrpc":"2.0","method":"cita_blockNumber","params":[],"id":74}'