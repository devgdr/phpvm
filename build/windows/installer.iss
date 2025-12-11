[Setup]
AppName=PHPVM
AppVersion=1.0.3
; A random GUID to identify this app
AppId={{D4166E7C-1B05-4B80-9285-802F3A79E234}
DefaultDirName={autopf}\PHPVM
DefaultGroupName=PHPVM
OutputBaseFilename=phpvm-setup
Compression=lzma
SolidCompression=yes
ArchitecturesInstallIn64BitMode=x64
PrivilegesRequired=admin
ChangesEnvironment=yes

[Tasks]
Name: "envPath"; Description: "Add to system PATH environment variable"; GroupDescription: "Additional icons:"; Flags: checkedonce

[Files]
; The Github Action will place the binary in ..\..\bin
Source: "..\..\bin\phpvm-windows-amd64.exe"; DestDir: "{app}"; DestName: "phpvm.exe"; Flags: ignoreversion

[Registry]
Root: HKLM; Subkey: "SYSTEM\CurrentControlSet\Control\Session Manager\Environment"; \
    ValueType: expandsz; ValueName: "Path"; ValueData: "{olddata};{app}"; \
    Check: NeedsAddPath('{app}'); Flags: preservestringtype

[Code]
function NeedsAddPath(Param: string): boolean;
var
  OrigPath: string;
begin
  if not RegQueryStringValue(HKEY_LOCAL_MACHINE,
    'SYSTEM\CurrentControlSet\Control\Session Manager\Environment',
    'Path', OrigPath)
  then begin
    Result := True;
    exit;
  end;
  // look for the path with leading and trailing semicolon
  // Pos() returns 0 if not found
  Result := Pos(';' + Param + ';', ';' + OrigPath + ';') = 0;
end;
