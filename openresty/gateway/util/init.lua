local cjson = require "cjson"

local _M = {}

local function split(str, pat)
   local t = {}  -- NOTE: use {n = 0} in Lua-5.0
   local fpat = "(.-)" .. pat
   local last_end = 1
   local s, e, cap = str:find(fpat, 1)
   while s do
      if s ~= 1 or cap ~= "" then
         table.insert(t,cap)
      end
      last_end = e+1
      s, e, cap = str:find(fpat, last_end)
   end
   if last_end <= #str then
      cap = str:sub(last_end)
      table.insert(t, cap)
   end
   return t
end
local function split_comma(str)
   return split(str,'[,]+')
end
local function split_path(str)
   return split(str,'[\\.]+')
end

function _M.xdomain(host)
	local names = split_path(host)
	local dname = "*"
	for i=1, #names do
		if i>1 then
	   		dname = dname .. "." .. names[i]
	   	end
	end
	return dname
end

function _M.getIngresNameByHost(host)
	local iname = ngx.shared["domains"]:get(host)
	if iname == nil then
		iname = ngx.shared["domains"]:get(_M.xdomain(host))
	end
	if iname == nil then
		return nil
	end
	return iname
end

function _M.debug()
	local doms = ngx.shared.domains
    ngx.log(ngx.ERR,"print domains")
	local keys = doms:get_keys()
    for i = 1, #keys do
        local d = doms:get(keys[i])
        ngx.log(ngx.ERR,"d=",keys[i],"i=",d)
    end
end

function _M.resetDomain()
	local ingress = ngx.shared["ingress"]
	local doms = ngx.shared.domains
	doms:flush_all()
 	local keys = ingress:get_keys() 
    for i = 1, #keys do
        local ingres = cjson.decode(ingress:get(keys[i]))
        local domains = split_comma(ingres.domain,",")
        for _, domain in ipairs(domains) do 
		    -- ngx.log(ngx.ERR,".domain=",domain,"ingres=",keys[i])
		    doms:set(domain, keys[i])
		end
    end
    -- _M.debug()
end

return _M