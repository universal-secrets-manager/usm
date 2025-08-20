package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Show audit log",
	Long:  "Show the audit log of secret operations.",
}

var auditShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show audit entries",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if audit log exists
		if _, err := os.Stat(".usm.audit.log"); os.IsNotExist(err) {
			fmt.Println("No audit log found")
			return
		}

		// Read the audit log
		data, err := os.ReadFile(".usm.audit.log")
		if err != nil {
			fmt.Printf("Failed to read audit log: %v\n", err)
			return
		}

		// Print the audit log
		fmt.Print(string(data))
	},
}

func init() {
	auditCmd.AddCommand(auditShowCmd)
	rootCmd.AddCommand(auditCmd)
}

// logAuditEntry logs an audit entry to the audit log file
func logAuditEntry(operation, key, profile string) {
	// Create or append to the audit log file
	f, err := os.OpenFile(".usm.audit.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		// If we can't log the audit entry, we shouldn't stop the operation
		return
	}
	defer f.Close()

	// Write the audit entry
	entry := fmt.Sprintf("[%s] %s %s in profile %s\n", time.Now().Format(time.RFC3339), operation, key, profile)
	f.WriteString(entry)
}