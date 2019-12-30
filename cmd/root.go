package cmd

import (
  "fmt"
  "github.com/sirupsen/logrus"
  "github.com/spf13/cobra"
  "github.com/thepwagner/archivist/index"
  "os"
  "time"

  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"
)


var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "archivist",
  Short: "Librarian for media horde",
  Long: `Tracks files across offline/removable and online media.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)
  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.archivist.yaml)")

  rootCmd.PersistentFlags().String( "index", "", "index file")
  viper.BindPFlag("index", rootCmd.PersistentFlags().Lookup("index"))
  viper.SetDefault("index", "./index.pb")
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    // Find home directory.
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".archivist" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".archivist")
  }

  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  logrus.SetLevel(logrus.DebugLevel)
  if err := viper.ReadInConfig(); err == nil {
    logrus.WithField("cfg", viper.ConfigFileUsed()).Debug("Using config file")
  }
}

func loadIndex() (*index.Index, error) {
  indexFile := viper.GetString("index")

  start := time.Now()
  idx, err := index.LoadIndex(indexFile)
  if err != nil {
    return nil, fmt.Errorf("loading index: %w", err)
  }
  logrus.WithFields(logrus.Fields{
    "path": indexFile,
    "dur": time.Since(start).Truncate(time.Millisecond).Seconds(),
  }).Debug("Loaded index")
  return idx, nil
}
