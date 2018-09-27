package databaseClient

type Cmd struct {
	err   error
	val []SetMember
}

func (cmd *Cmd) Result() ([]SetMember, error){
	return cmd.val, cmd.err
}

func (cmd *Cmd) Err() error {
	return cmd.err
}
