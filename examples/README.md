# Examples

## Description

This directory contains examples that are mostly used for documentation, but can also be run/tested manually via the Terraform CLI.

Subdirectory [./basic](./basic) shows how to create basic integrations of multiple building blocks to achieve everyday tasks.

Remaining subdirectories are focused on specific single building blocks.

## Document Generation Tool

The document generation tool looks for files in the following locations by default. All other *.tf files besides the ones mentioned below are ignored by the documentation tool. This is useful for creating examples that can run and/or ar testable even if some parts are not relevant for the documentation.

* **provider/provider.tf** example file for the provider index page
* **data-sources/`full data source name`/data-source.tf** example file for the named data source page
* **resources/`full resource name`/resource.tf** example file for the named data source page
