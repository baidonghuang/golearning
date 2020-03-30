local util = require "gateway.util"
local ngx_set_header = ngx.req.set_header
local cjson = require "cjson"

local _M = {}

function _M.exec()

    local ingresName = util.getIngresNameByHost(ngx.var.host)
    ngx.log(ngx.INFO, "host:" .. ngx.var.host, "ingress:" .. ngx.var.ingres,  "uri:" .. ngx.var.request_uri)


end

return _M