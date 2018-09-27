package redisClient

type Cmd struct {
	err   error
	val []Z
}

func (cmd *Cmd) Result() ([]Z, error){
	return cmd.val, cmd.err
}

func (cmd *Cmd) Err() error {
	return cmd.err
}
