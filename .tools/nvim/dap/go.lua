local dap = require("dap")

dap.configurations.go = {
  {
    type = "go",
    name = "Debug infra-api",
    request = "launch",
    program = vim.g.root_dir .. "/apps/infra",
    args = { "--dev" },
    console = "externalTerminal",
    -- externalTerminal = true,
    envFile = {
      vim.g.root_dir .. "/apps/infra" .. "/.secrets/env",
    },
  },
  {
    type = "go",
    name = "Debug console-api",
    request = "launch",
    program = vim.g.root_dir .. "/apps/console",
    args = { "--dev" },
    console = "externalTerminal",
    -- externalTerminal = true,
    envFile = {
      vim.g.root_dir .. "/apps/console" .. "/.secrets/env",
    },
    -- dlvToolPath =
    -- "/usr/local/go/bin/dlv --headless=true --api-version=2 -r stdout:/tmp/debug.stdout -r stderr:/tmp/debug2.stderr",
  },
  {
    type = "go",
    name = "Debug finance-api",
    request = "launch",
    program = vim.g.root_dir .. "/apps/finance",
    args = { "--dev" },
    console = "externalTerminal",
    -- externalTerminal = true,
    envFile = {
      vim.g.root_dir .. "/apps/finance" .. "/.secrets/env",
    },
  },
  {
    type = "go",
    name = "Debug iam-api",
    request = "launch",
    program = vim.g.root_dir .. "/apps/iam",
    args = { "--dev" },
    console = "externalTerminal",
    -- externalTerminal = true,
    envFile = {
      vim.g.root_dir .. "/apps/iam" .. "/.secrets/env",
    },
  },
  {
    type = "go",
    name = "Debug message-office-api",
    request = "launch",
    program = vim.g.root_dir .. "/apps/message-office",
    args = { "--dev" },
    console = "externalTerminal",
    -- externalTerminal = true,
    envFile = {
      vim.g.root_dir .. "/apps/message-office" .. "/.secrets/env",
    },
  },
  {
    type = "go",
    name = "Debug container-registry-api",
    request = "launch",
    program = vim.g.root_dir .. "/apps/container-registry",
    args = { "--dev" },
    console = "externalTerminal",
    -- externalTerminal = true,
    envFile = {
      vim.g.root_dir .. "/apps/container-registry" .. "/.secrets/env",
    },
  },
}
