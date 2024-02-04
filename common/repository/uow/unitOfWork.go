package uow

import "context"

type UnitOfWork struct {
	commands []Command
}

func (u *UnitOfWork) AddCommand(command Command) {
	u.commands = append(u.commands, command)
}

func (u *UnitOfWork) Commit(ctx context.Context) error {
	for _, command := range u.commands {
		if err := command.Execute(ctx); err != nil {
			return err
		}
	}
	u.commands = []Command{} // Clear the commands after successful execution
	return nil
}

func (u *UnitOfWork) Rollback() {
	u.commands = []Command{} // Simply clear the commands to rollback
}
