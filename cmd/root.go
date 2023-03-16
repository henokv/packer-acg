/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/henokv/packer-acg/internal"
	"github.com/spf13/cobra"
	"go.mercari.io/hcledit"
	"os"
	"time"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "packer-acg",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },,
	Run: rootCmdRun,
}

func rootCmdRun(cmd *cobra.Command, args []string) {

	inPath, err := cmd.Flags().GetString("file")
	cobra.CheckErr(err)

	outPath, err := cmd.Flags().GetString("output-file")
	//cobra.
	if err != nil {
		if err.Error() == "flag accessed but not defined: output-file" {
			outPath = inPath
		} else {
			cobra.CheckErr(err)
		}
	}

	editor, err := hcledit.ReadFile(inPath)
	cobra.CheckErr(err)

	cliLogin, err := cmd.Flags().GetBool("azure-cli")
	cobra.CheckErr(err)
	if cliLogin {
		internal.SetAzureCLIAuth(editor)
	}

	dryRun, err := cmd.Flags().GetBool("dry-run")
	cobra.CheckErr(err)

	azureImageGalleryName, err := cmd.Flags().GetString("gallery-name")
	cobra.CheckErr(err)
	azureImageGalleryResourceGroupName, err := cmd.Flags().GetString("gallery-rg-name")
	cobra.CheckErr(err)
	stType, err := cmd.Flags().GetString("storage-type")
	cobra.CheckErr(err)
	name, err := cmd.Flags().GetString("image-name")
	cobra.CheckErr(err)
	location, err := cmd.Flags().GetString("location")
	cobra.CheckErr(err)
	if name != "" && azureImageGalleryName != "" && azureImageGalleryResourceGroupName != "" {
		internal.SetAzureImageGallery(editor)
		internal.CreateVariable(editor, "gallery_name", "string", azureImageGalleryName)
		internal.CreateVariable(editor, "gallery_image_name", "string", name)
		internal.CreateVariable(editor, "image_gallery_resource_group", "string", azureImageGalleryResourceGroupName)
		internal.CreateVariable(editor, "storage_account_type", "string", stType)
		internal.CreateVariable(editor, "replication_regions", "list(string)", "[]")
	}

	internal.SetVariableDefailt(editor, "location", location)
	instanceNumber, err := cmd.Flags().GetInt("instance-number")
	cobra.CheckErr(err)
	version := fmt.Sprintf("%04d.%02d.%02d%02d", time.Now().Year(), time.Now().Month(), time.Now().Day(), instanceNumber)
	internal.SetVariableDefailt(editor, "image_version", version)

	if dryRun {
		editor.Write(os.Stdout)
	} else {
		editor.WriteFile(outPath)
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("azure-cli", "c", true, "Should Azure cli be used for authentication")
	rootCmd.Flags().StringP("gallery-name", "n", "", "Azure image gallery name")
	rootCmd.Flags().StringP("gallery-rg-name", "g", "", "Azure image gallery resource group name")
	rootCmd.Flags().String("storage-type", "Premium_LRS", "Azure storage type")
	rootCmd.Flags().StringP("image-name", "i", "", "Image name")
	rootCmd.Flags().StringP("location", "l", "westeurope", "Azure location")
	rootCmd.Flags().Int("instance-number", 1, "Used to set image version 2023.16.0301. Last two digits are instance number")
	rootCmd.MarkFlagsRequiredTogether("gallery-name", "gallery-rg-name", "image-name")

	rootCmd.Flags().BoolP("dry-run", "d", false, "Should it cat the ouput")
	rootCmd.Flags().StringP("file", "f", "", "The input file")
	rootCmd.Flags().StringP("output-file", "o", "", "The output file, if not provided uses the input file")
	rootCmd.MarkFlagRequired("file")

}
