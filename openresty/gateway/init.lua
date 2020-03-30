local rewrite = require "gateway.rewrite"


local _Instance = {}

function _Instance.init()
    print("--------------> hello init")
end

function _Instance.init_worker(ctx)
    if ctx.count == nil then
        ctx.count = 1
    else
        ctx.count = ctx.count + 1
    end
    local ingress = ngx.shared["userdict"]
    ingress:set("init_worker", "init_worker:" .. ctx.count)
end

function _Instance.balancer(ctx)
    if ctx.count == nil then
        ctx.count = 1
    else
        ctx.count = ctx.count + 1
    end
    local ingress = ngx.shared["userdict"]
    ingress:set("balancer", "balancer:" .. ctx.count)
end

-- 1
function _Instance.rewrite(ctx)
    rewrite.exec()
end

-- 2
function _Instance.access(ctx)
    if ctx.count == nil then
        ctx.count = 1
    else
        ctx.count = ctx.count + 1
    end
    local ingress = ngx.shared["userdict"]
    ingress:set("access", "access:" .. ctx.count)
end

-- 3
function _Instance.content(ctx)
    if ctx.count == nil then
        ctx.count = 1
    else
        ctx.count = ctx.count + 1
    end
    local ingress = ngx.shared["userdict"]
    ingress:set("content","content:" ..  ctx.count)
    local rewrite, flags = ingress:get("rewrite")
    local access, flags = ingress:get("access")
    local header, flags = ingress:get("header_filter")
    local body, flags = ingress:get("body_filter")
    local init_worker, flags = ingress:get("init_worker")
    local balancer, flags = ingress:get("balancer")
    local content, flags = ingress:get("content")
    ngx.say(rewrite, access, header, body, balancer, init_worker, content)
end

-- 4
function _Instance.header_filter(ctx)
    if ctx.count == nil then
        ctx.count = 1
    else
        ctx.count = ctx.count + 1
    end
    local ingress = ngx.shared["userdict"]
    ingress:set("header_filter", "header_filter:" .. ctx.count)
end

-- 6
function _Instance.body_filter(ctx)
    if ctx.count == nil then
        ctx.count = 1
    else
        ctx.count = ctx.count + 1
    end
    local ingress = ngx.shared["userdict"]
    ingress:set("body_filter", "body_filter:" .. ctx.count)
end

function _Instance.handle_error()

    local ingress = ngx.shared["userdict"]
    local rewrite, flags = ingress:get("rewrite")
    local access, flags = ingress:get("access")
    local header, flags = ingress:get("header_filter")
    local body, flags = ingress:get("body_filter")
    local robots, flags = ingress:get("robots")
    local init_worker, flags = ingress:get("init_worker")
    local balancer, flags = ingress:get("balancer")
    local content, flags = ingress:get("content")
    local errmsg = ngx.var.api_gw_error
    ngx.say(errmsg, rewrite, access, header, body, robots, balancer, init_worker, content)
    print("--------------> hbd handle_error  <----------------")

end

function _Instance.robots()
    local ingress = ngx.shared["userdict"]
    ingress:set("robots", "hello robots")
    ngx.var.api_gw_error = "---------> failed to set the current peer address: "
    return ngx.exit(404)
end

return _Instance