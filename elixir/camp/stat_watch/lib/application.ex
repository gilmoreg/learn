defmodule StatWatch.Application do
  use Application

  def start(_type, _args) do
    import Supervisor.Spec, warn: false

    children = [
      worker(StatWatch.Scheduler, [])
    ]

    options = [strategy: :one_for_one, name: StatWatch.Supervisor]
    Supervisor.start_link(children, options)
  end
end
