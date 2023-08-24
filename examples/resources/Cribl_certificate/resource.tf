resource "Cribl_certificate" "my_certificate" {
  ca          = "...my_ca..."
  cert        = "...my_cert..."
  description = "...my_description..."
  id          = "89bd9d8d-69a6-474e-8f46-7cc8796ed151"
  in_use = [
    "...",
  ]
  passphrase = "...my_passphrase..."
  priv_key   = "...my_priv_key..."
}