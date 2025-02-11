package config

type SshConfig struct {
	Host string
	Port string
}

func GetSshConfig() SshConfig {
	host := readEnvWithFallback("SANGAM_QUIZ_SSH_HOST", "localhost")
	port := readEnvWithFallback("SANGAM_QUIZ_SSH_PORT", "23235")
	return SshConfig{Host: host, Port: port}
}
