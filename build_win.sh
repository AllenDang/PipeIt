echo "
<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>
<assembly xmlns=\"urn:schemas-microsoft-com:asm.v1\" manifestVersion=\"1.0\">
    <assemblyIdentity version=\"1.0.0.0\" processorArchitecture=\"*\" name=\"PipeIt\" type=\"win32\"/>
    <dependency>
        <dependentAssembly>
            <assemblyIdentity type=\"win32\" name=\"Microsoft.Windows.Common-Controls\" version=\"6.0.0.0\" processorArchitecture=\"*\" publicKeyToken=\"6595b64144ccf1df\" language=\"*\"/>
        </dependentAssembly>
    </dependency>
    <application xmlns=\"urn:schemas-microsoft-com:asm.v3\">
        <windowsSettings>
            <dpiAwareness xmlns=\"http://schemas.microsoft.com/SMI/2016/WindowsSettings\">PerMonitorV2, PerMonitor</dpiAwareness>
            <dpiAware xmlns=\"http://schemas.microsoft.com/SMI/2005/WindowsSettings\">True</dpiAware>
        </windowsSettings>
    </application>
</assembly>
" > PipeIt.exe.manifest

rsrc -manifest PipeIt.exe.manifest -ico ./resource/pipeline.ico -arch amd64 -o rsrc.syso

go build -ldflags='-s -w -H windowsgui -linkmode external -extldflags -static' .

rm PipeIt.exe.manifest
rm rsrc.syso
