local function i(value) 
  print(vim.inspect(value))
end

print("hhhh")
cur = vim.api.nvim_buf_get_name(0)
print("current buffer" .. cur)


local language_tree = vim.treesitter.get_parser(0, "go")
local q = require"vim.treesitter.query" 

local syntax_tree = language_tree:parse()

local root = syntax_tree[1]:root()


local query = vim.treesitter.parse_query('go', [[
(function_declaration 
  name: (identifier) @test) 
]])

for _, captures, metadata in query:iter_matches(root, 0) do 
  i(q.get_node_text(captures[1], 0))
end
