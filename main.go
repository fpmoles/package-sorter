package main

import (
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/fpmoles/package-sorter/internal/sorter"
)

func main() {
	var debug bool
	var width, height, length, mass float64

	rootCmd := &cobra.Command{
		Use:   "package-sorter [--width W] [--height H] [--length L] [--mass M] | <width> <height> <length> <mass>",
		Short: "Dispatch a package to the correct stack based on its dimensions and mass",
		Args:  cobra.MaximumNArgs(4),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if debug {
				log.SetLevel(log.DebugLevel)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// If positional args provided, they override flags (in order: width height length mass)
			if len(args) > 0 {
				if len(args) != 4 {
					return fmt.Errorf("expected 4 positional arguments (width height length mass), got %d", len(args))
				}
				names := []string{"width", "height", "length", "mass"}
				vals := []*float64{&width, &height, &length, &mass}
				for i, arg := range args {
					v, err := strconv.ParseFloat(arg, 64)
					if err != nil {
						return fmt.Errorf("invalid %s: %q", names[i], arg)
					}
					*vals[i] = v
				}
			}

			log.WithFields(log.Fields{
				"width":  width,
				"height": height,
				"length": length,
				"mass":   mass,
			}).Debug("sorting package")

			stack, err := sorter.Sort(width, height, length, mass)
			if err != nil {
				return err
			}

			log.WithField("stack", stack).Debug("dispatch decision")
			fmt.Println(stack)
			return nil
		},
	}

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug logging")
	rootCmd.Flags().Float64Var(&width, "width", 0, "width in cm")
	rootCmd.Flags().Float64Var(&height, "height", 0, "height in cm")
	rootCmd.Flags().Float64Var(&length, "length", 0, "length in cm")
	rootCmd.Flags().Float64Var(&mass, "mass", 0, "mass in kg")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
