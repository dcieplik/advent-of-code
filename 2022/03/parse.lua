local function i(value) 
  print(vim.inspect(value))
end

local getPath=function(str,sep)
    sep=sep or'/'
    return str:match("(.*"..sep..")")
end

local q = require"vim.treesitter" 
local ts_utils = require"nvim-treesitter.ts_utils" 

function get_current_function_name() 
  local current_node = ts_utils.get_node_at_cursor() 
  if not current_node then return "" end
  
  local expr = current_node

  while expr do
    parent = expr:parent()
    if expr:type() == 'function_declaration' and parent:type() == "source_file" then
        break
    end
    expr = expr:parent()
  end

  if not expr then return "" end

  return (ts_utils.get_node_text(expr:child(1)))[1]
end

function exec() 
  local name = get_current_function_name()
  local cur = vim.api.nvim_buf_get_name(0)
  local path = getPath(cur)
  local handle = io.popen("cd " .. path .. "; go test -run " .. name .. " -v")
  local result = handle:read("*a")
  handle:close()
  print(result)
end
