---
title: Parboiled简介
tags:
---

# Parboiled ==
```
#!div class=important
'''[[Color(,red,真令人失望)]]:'''
Parboiled 的实现中使用了一些Dalvik还不支持的JVM特性(还没搞清楚，是和字节码相关的)， 我发现即便将整个源码都编译到apk中
依然无法让Parboiled正常工作，它的几个核心类运行期无法生成。
```
http://wiki.github.com/sirthias/parboiled/
 
 * Apache2 Lisense
 * No need to write intermediate syntax description script like yacc or javacc ..
 * Easy to use(PEG)
# ImapParser
We can create a strict imap parser by the FormalSyntax of IMAP4, here is a part of imap paser.
```
#!java
package client.protocol.imap;
import org.parboiled.BaseParser;
import org.parboiled.Rule;
public class ImapParser extends BaseParser<Object>{
        public Rule Expression() {
                // TODO Auto-generated method stub
                return null;
        }
        /**
         * atom-specials   = "(" / ")" / "{" / SP / CTL / list-wildcards /quoted-specials / resp-specials
         * TODO: what's CTL ???
         */
        public Rule AtomSpecials() {
                return FirstOf("(",")","{",SP(),/*CTL(),*/ListWildcards(), QuotedSpecials(), RespSpecials());
        }
    
        /**
         * list-wildcards  = "%" / "*"
         */
        public Rule ListWildcards() {
                return FirstOf("%","*");
        }
        /**
         * quoted-specials = DQUOTE / ""
         */
        public Rule QuotedSpecials() {
                return FirstOf(DQUOTE(), '\');
        }
        /**
         * quoted          = DQUOTE *QUOTED-CHAR DQUOTE
         */
        public Rule Quoted() {
                return Sequence(DQUOTE(), ZeroOrMore(QuotedChar()),DQUOTE());
        }
        /*
         * <any TEXT-CHAR except quoted-specials> / "" quoted-specials
         * @return
         */
        public Rule QuotedChar() {
                return FirstOf(Sequence("\",QuotedSpecials()), Test(QuotedSpecials()), TEXT_CHAR());
        }
        /**
         * TEXT-CHAR       = <any CHAR except CR and LF>
         * TODO: CHAR is CHAR8 ??
         */
        public Rule TEXT_CHAR() {
                return FirstOf(Test(CR()),Test(LF()),CHAR8());
        }
        /**
         * Unsigned 32-bit integer
         * (0 <= n < 4,294,967,296)
         * @return
         */
        public Rule Number() {
                return OneOrMore(Digit());
        }
        /*
         * quoted / literal
         */
        public Rule String() {
                return FirstOf(Quoted(),Literal());
        }
        /*
         * "{" number "}" CRLF *CHAR8
         */
        public Rule Literal() {
                return Sequence("{",Number(),"}",CRLF(),ZeroOrMore(CHAR8()));
        }
        /**
         * ASTRING-CHAR   = ATOM-CHAR / resp-specials
         */
        public Rule ASTRING_CHAR() {
                return FirstOf(ATOM_CHAR(), RespSpecials());
        }
        /**
         * <any CHAR except atom-specials> 
         * @return 
         * TODO:??? CHAR is CHAR8 ???
         */
        public Rule ATOM_CHAR() {
                return FirstOf(Test(AtomSpecials()), CHAR8());
        }
        /**
         * astring         = 1*ASTRING-CHAR / string
         * @return
         */
        public Rule AString() {
                return FirstOf(OneOrMore(AString()), String());
        }
        //---------------------------------------------------------[ Single ]
        public Rule RespSpecials() {
                return Ch(']');
        }
        /**
         * CHAR8           = %x01-ff
         */
        public Rule CHAR8() {
                return CharRange('', 'ÿ');
        }
        public Rule Digit() {
                return CharRange('0', '9');
        }
        public Rule SP() {
                return Ch(' ');
        }
        public Rule CR() {
                return Ch('');
        }
        public Rule LF() {
                return Ch('
');
        }
        public Rule CRLF() {
                return Sequence(CR(),LF());
        }
        public Rule DQUOTE() {
                return Ch('"');
        }
}
```
For a simple parse the Number defined by Imap4:
```
#!java
                InputStreamReader is = new InputStreamReader(System.in);
                BufferedReader    br = new BufferedReader(is);
                boolean         quit = false;
                ImapParser parser = Parboiled.createParser(ImapParser.class);
                while(!quit) {
                        String line = "";;
                        try {
                                line = br.readLine();
                        } catch (IOException e) {
                                // TODO Auto-generated catch block
                                e.printStackTrace();
                        }
                        ParsingResult<?> result = ReportingParseRunner.run(parser.Number(), line);
                        String parseTreePrintOut = ParseTreeUtils.printNodeTree(result); 
                        System.out.println(parseTreePrintOut);
                }
```
```
12345 <== enter 
[Number] '12345'
    [Digit] '1'
    [Digit] '2'
    [Digit] '3'
    [Digit] '4'
    [Digit] '5'
abc <== enter
<Noting Happened>
```
