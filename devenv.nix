{ pkgs, ... }:

{

  dotenv.enable = true;

  languages.go = {
    enable = true;
    delve.enable = true;
    lsp.enable = true;
    version = "1.27.1";
  };

  packages = with pkgs; [
    gosec
    google-cloud-sdk
  ];

  tasks."boot-dev:submit" = {
    exec = ''yes | bootdev run -s'';
  };

}
