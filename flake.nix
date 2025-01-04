{
  description = "gitjika";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      nixpkgs,
      utils,
      ...
    }:
    utils.lib.eachDefaultSystem (
      system:
      let
        overlays = [
          (self: super: {
            templ = super.templ.overrideAttrs (old: rec {
              version = "0.3.819";
              src = pkgs.fetchFromGitHub {
                owner = "a-h";
                repo = "templ";
                rev = "v${version}";
                hash = "sha256-kTP/DLnou3KETZRtvHdeiMmRW6xldgZBAn9O9p9s/MA=";
              };
              vendorHash = "sha256-ipLn52MsgX7KQOJixYcwMR9TCeHz55kQQ7fgkIgnu7w=";
            });
          })
        ];
        pkgs = import nixpkgs {
          inherit system overlays;
        };
      in
      {
        devShell =
          with pkgs;
          mkShell {
            buildInputs = [
              go
              gopls
              golangci-lint
              templ
              tailwindcss
            ];
          };
      }
    );
}
