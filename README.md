### Operation guide:

1. add -javaagent:/path/to/ja-netfilter.jar=jetbrains to your vmoptions (manual or auto)
2. log out of the jb account in the 'Licenses' window
3. use command to build the user cert
4. use command to generate a custom license
5. don't care about the activation time, it is a fallback license and will not expire

#### Enjoy it~

##### JBR17:

> for Java version 17+, you need add these 2 lines to your vmoptions file: (for manual, without any whitespace chars)

```vmoptions
--add-opens=java.base/jdk.internal.org.objectweb.asm=ALL-UNNAMED
--add-opens=java.base/jdk.internal.org.objectweb.asm.tree=ALL-UNNAMED
```