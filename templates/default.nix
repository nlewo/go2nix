{ stdenv, lib, go, goPackages, fetchgit, fetchhg, fetchbzr, fetchsvn }:

goPackages.buildGoPackage rec {
  name = "{{ .Name }}-${stdenv.lib.strings.substring 0 7 rev}";
  rev = "{{ .Revision }}";

  buildInputs = [ go ];
  goPackagePath = "{{ .ImportPath }}";

  src = fetchgit {
    url = "{{ .VcsRepo }}";
    rev = "{{ .Revision }}";
    sha256 = "{{ .Hash }}";
  };

  extraSrcs = import ./deps.nix {
    inherit fetchgit fetchhg fetchbzr fetchsvn;
  };
}
