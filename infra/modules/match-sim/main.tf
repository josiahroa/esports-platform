module "ecr_event_stream" {
  source  = "terraform-aws-modules/ecr/aws"
  version = "2.0.0"
}

module "lambda_event_stream" {
  source = "terraform-aws-modules/lambda/aws"

  function_name = "event-stream"
  description   = "Event stream for match simulation events"

  create_package = false

  image_uri    = "${module.ecr_event_stream.repository_url}:latest"
  package_type = "Image"
}
