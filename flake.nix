{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = { self, nixpkgs, flake-utils, treefmt-nix }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };

        treefmtEval = treefmt-nix.lib.evalModule pkgs {
          projectRootFile = "flake.nix";
          programs = {
            nixpkgs-fmt.enable = true;
            gofmt.enable = true;
          };
        };

        sangam-quiz = pkgs.buildGoModule {
          name = "sangam-quiz";
          version = "0.0.1";
          #vendorHash = pkgs.lib.fakeHash;
          vendorHash = "sha256-l5i1C7SFbAQhtOrVRXuvY7S7T/oKlxejrVgBj1hP+Pw=";
          src = ./.;
          subPackages = "cmd/sangam-quiz";
        };

        sangam-quiz-ssh = pkgs.buildGoModule {
          name = "sangam-quiz-ssh";
          version = "0.0.1";
          #vendorHash = pkgs.lib.fakeHash;
          vendorHash = "sha256-l5i1C7SFbAQhtOrVRXuvY7S7T/oKlxejrVgBj1hP+Pw=";
          src = ./.;
          subPackages = "cmd/sangam-quiz-ssh";
        };
      in
      {
        formatter = treefmtEval.config.build.wrapper;
        checks.formatter = treefmtEval.config.build.check self;

        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
          ];
        };

        packages = flake-utils.lib.flattenTree {
          default = sangam-quiz;
          ssh = sangam-quiz-ssh;
        };
      }
    );
}
