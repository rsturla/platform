package command

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "Set the log level (trace, debug, info, warn, error, fatal, panic)")
	rootCmd.AddCommand(versionCmd)
}
