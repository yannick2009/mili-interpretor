{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";

  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };

        name = "Mili interpretor";
        version = "1.0.0";

        devPkgs = with pkgs; [
          go
          tinygo
          gopls
          just
        ];
      in
      {
        devShells = {
          default = pkgs.mkShell {
            buildInputs = devPkgs;
            shellHook = ''
              echo "building ${name} - v${version}"
            '';
          };
        };
        # Create a new package
        packages = { };
      }
    );
}
