package internal

import "go.mercari.io/hcledit"

func CreateVariable(editor hcledit.HCLEditor, name string, valType string, value string) {
	editor.Create("variable", hcledit.BlockVal(name))
	editor.Create("variable."+name+".type", hcledit.RawVal(valType))
	editor.Create("variable."+name+".default", value)
}

func SetVariableDefailt(editor hcledit.HCLEditor, name string, defaultValue string) {
	editor.Create("variable."+name+".default", defaultValue)
}

func SetAzureCLIAuth(editor hcledit.HCLEditor) {
	editor.Create("source.azure-arm.*.use_azure_cli_auth", true)
	editor.Delete("source.azure-arm.*.client_id")
	editor.Delete("source.azure-arm.*.client_secret")
	editor.Delete("source.azure-arm.*.client_cert_path")
}

func SetAzureImageGallery(editor hcledit.HCLEditor) {
	editor.Create("source.azure-arm.*.shared_image_gallery_destination", hcledit.BlockVal())
	editor.Create("source.azure-arm.*.shared_image_gallery_destination.gallery_name", hcledit.RawVal("\"${var.gallery_name}\""))
	editor.Create("source.azure-arm.*.shared_image_gallery_destination.image_name", hcledit.RawVal("\"${var.gallery_image_name}\""))
	editor.Create("source.azure-arm.*.shared_image_gallery_destination.image_version", hcledit.RawVal("\"${var.image_version}\""))
	editor.Create("source.azure-arm.*.shared_image_gallery_destination.resource_group", hcledit.RawVal("\"${var.image_gallery_resource_group}\""))
	editor.Create("source.azure-arm.*.shared_image_gallery_destination.subscription", hcledit.RawVal("\"${var.subscription_id}\""))
	editor.Create("source.azure-arm.*.shared_image_gallery_destination.replication_regions", hcledit.RawVal("\"${var.replication_regions}\""))
	editor.Create("source.azure-arm.*.shared_image_gallery_destination.storage_account_type", hcledit.RawVal("\"${var.storage_account_type}\""))

	editor.Delete("source.azure-arm.*.capture_container_name")
	editor.Delete("source.azure-arm.*.capture_name_prefix")
	editor.Delete("source.azure-arm.*.resource_group_name")
	editor.Delete("source.azure-arm.*.storage_account")
}
