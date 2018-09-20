---
title: emacs
tags:
---
<!-- toc -->
[[TOC]]
# AutoCompleteMode =
# 6. Source ==
Sourceç¨æ¥æä»£è¡¥å¨ä¿¡æ¯æ¥æºï¼ AutoCompleteModeéè¿Sourceè¿ä¸ªæ¦å¿µæ¥æè¿°è¡¥å¨ä¿¡æ¯çæ¥æºã
Source is a concept that insures a extensibility of auto-complete-mode. Simply saying, source is a description about:
ä½ å¯ä»¥éè¿å®ä¹èªå·±çSourceæ¥è¾¾ä¸°å¯è¡¥å¨ä¿¡æ¯ãæ¯ç§Sourceå¿é¡»ä»¥`ac-source-`å¼å¤´ã
How to generate completion candidates
How to complete
How to show
 
Anybody who know about Emacs Lisp a little can define a source easily. See extend for how to define a source. Here we can explain how to use builtin sources.
Usually a name of source starts with ac-source-. So you can list up sources with apropos (M-x apropos RET ^ac-source-). You may see ac-source-filename and ac-source-dictionary which are entities of sources.
å½ä½ å·²ç»å¼å¯äº AutoCompleteMode, ä½ å¯ä»¥å¨*scratch*ä¸­å¯¹`ac-sources`æ±å¼ï¼è¿æ ·ä½ å°±å¯ä»¥çå°å½åæ¯æåªäºSourceäº:
```
ac-sources C-x-e
```
ä¹å¯ä»¥éè¿C-h v, æ¥çac-sourcesçå¼ã
 
# 6.1. Using Source ===
å¨ä¸åæ¨¡å¼ä¸ï¼éç½®ä¸åçac-sources, æ¯å¦:
æä»¬å¸æå¨lisp-modeæ¶åªä½¿ç¨ä»¥ä¸ä¸¤ä¸ªSource:
  * ac-source-symbols 
  * ac-source-words-in-same-mode-buffers
å¯ä»¥è¿ä¹å:
```
#!el
(defun my-ac-emacs-lisp-mode ()
  (setq ac-sources '(ac-source-symbols ac-source-words-in-same-mode-buffers)))
(add-hook 'emacs-lisp-mode-hook 'my-ac-emacs-lisp-mode)
```
é¢å®ä¹Sourceä¸è§:
||=é¢å®ä¹Sourceå=||=è¯´æ=||
|| ac-source-yasnippet || A source for Yasnippet to complete and expand snippets.                                          ||
|| ac-source-dictionary|| A source for dictionary. See completion by dictionary about dictionary.                          ||
|| ac-source-eclim     || A source for EmacsEclim.                                                                         ||                                             
|| ac-source-filename  || A source for completing file name. Completion will be started after inserting /.                 ||
|| ac-source-files-in-current-dir || A source for completing files in a current directory. It may be useful with eshell.   ||
|| ac-source-functions || A source for completing Emacs Lisp functions. It is available only after (.                      ||
|| ac-source-gtags     || A source for completing tags of Global.                                                          || 
|| ac-source-semantic  || A source for Semantic. It can be used for completing member name for C/C++.                      ||
|| ac-source-symbols   || A source for completing Emacs Lisp symbols.                                                      ||
|| ac-source-variables || A source for completing Emacs Lisp variables                                                     ||
|| ac-source-words-in-all-buffer || A source for completing words in all buffer. Unlikely ac-source-words-in-same-mode-buffers, it doesn't regard major-mode. ||
|| ac-source-words-in-buffer || A source for completing words in a current buffer.                                         ||
|| ac-source-words-in-same-mode-buffers ||  source for completing words which are collected over buffers whom major-mode is same to of a current buffer. ||
|| ac-source-css-property || A source for CSS properties                                                                   ||
# 7. Tips ==
# 7.1. Not to complete automatically ===
If you are being annoyed with displaying completion menu, you can disable automatic starting completion by setting ac-auto-start to nil.
```
(setq ac-auto-start nil)
```
You need to bind some key to auto-complete command (because you need to complete anyway). For example, bind to ac-mode-map, which is a key map for auto-complete-mode enabled buffer:
```
(define-key ac-mode-map (kbd "M-TAB") 'auto-complete)
```
Or bind to global key map.
```
(global-set-key "¯" 'auto-complete)
```
In addition, if you allow to start completion automatically but also want to be silent as much as possible, you can do it by setting ac-auto-start to an prefix length integer. For example, if you want to start completion automatically when you has inserted 4 or more characters, just set ac-auto-start to 4:
```
(setq ac-auto-start 4)
```
Setting ac-auto-start to large number will result in good for performance. Lesser ac-auto-start, more high cost to produce completion candidates, because there will be so many candidates necessarily. If you feel auto-complete-mode is stalling, change ac-auto-start to a larger number or nil.
See ac-auto-start for more details.
And consider to use trigger key.
# 7.2. Not to show completion menu automatically ===
There is another approach to solve the annoying problem is that not to show completion menu automatically. Not to show completion menu automatically, set ac-auto-show-menu to nil.
```
(setq ac-auto-show-menu nil)
```
When you select or filter candidates, completion menu will be shown.
In other way, you can delay showing completion menu by setting ac-auto-show-menu to seconds in real number.
```
;; Show 0.8 second later
(setq ac-auto-show-menu 0.8)
```
This interface has both good points of completely automatic completion and completely non-automatic completion. This may be default in the future.
# 7.3. Stop completion ===
You can stop completion by pressing C-g. However you won't press C-g while defining a macro. In such case, it is a good idea to bind some key to ac-completing-map.
```
(define-key ac-completing-map "¯" 'ac-stop)
```
Now you can stop completion by pressing M-/.
# 7.4. Finish completion by TAB ===
As we described above, there is many behaviors in TAB. You need to use TAB and RET properly, but there is a simple interface that bind RET to original and TAB to finish completion:
```
(define-key ac-completing-map "	" 'ac-complete)
(define-key ac-completing-map "" nil)
```
# 7.5. Select candidates with C-n/C-p only when completion menu is displayed ===
By evaluating the following code, you can select candidates with C-n/C-p, but it might be annoying sometimes.
```
;; Bad config
(define-key ac-completing-map "" 'ac-next)
(define-key ac-completing-map "" 'ac-previous)
```
In this case, it is better that selecting candidates is enabled only when completion menu is displayed so that the key input will not be taken as much as possible. ac-menu-map is a keymap for completion on completion menu which is enabled when ac-use-menu-map is t.
```
(setq ac-use-menu-map t)
;; Default settings
(define-key ac-menu-map "" 'ac-next)
(define-key ac-menu-map "" 'ac-previous)
```
See ac-use-menu-map and ac-menu-map for more details.
# 7.6. Not to use quick help ===
A tooltip help that is shown when completing is called quick help. You can disable it if you don't want to use it:
```
(setq ac-use-quick-help nil)
```
# 7.7. Change a height of completion menu ===
Set ac-menu-height to number of lines.
```
;; 20 lines
(setq ac-menu-height 20)
```
# 7.8. Enable auto-complete-mode automatically for specific modes ===
auto-complete-mode won't be enabled automatically for modes that are not in ac-modes. So you need to set if necessary:
```
(add-to-list 'ac-modes 'brandnew-mode)
```
# 7.9. Ignore case ===
There is three ways to distinguish upper case and lower case.
```
;; Just ignore case
(setq ac-ignore-case t)
;; Ignore case if completion target string doesn't include upper characters
(setq ac-ignore-case 'smart)
;; Distinguish case
(setq ac-ignore-case nil)
```
Default is smart.
# 7.10. Stop completion automatically after inserting specific words ===
Set ac-ignores to words that stops completion automatically. In ruby, some people want to stop completion automatically after inserting "end":
```
(add-hook 'ruby-mode-hook
          (lambda ()
            (make-local-variable 'ac-ignores)
            (add-to-list 'ac-ignores "end")))
```
Note that ac-ignores is not a buffer local variable, so you need to make it buffer local with make-local-variable if it is buffer specific setting.
# 7.11. Change colors ===
Colors settings are following:
||=Face=||=Description=||
||ac-completion-face||Foreground color of inline completion||
||ac-candidate-face ||Color of completion menu
||ac-selection-face ||Selection color of completion menu
To change face background color, use set-face-background. To change face foreground color, use set-face-foreground. To set underline, use set-face-underline.
```
;; Examples
(set-face-background 'ac-candidate-face "lightgray")
(set-face-underline 'ac-candidate-face "darkgray")
(set-face-background 'ac-selection-face "steelblue")
```
# 7.12. Change default sources ===
Read source first if you don't familiar with sources. To change default of sources, use setq-default:
```
(setq-default ac-sources '(ac-source-words-in-all-buffer))
```
# 7.13. Change sources for specific major modes ===
For example, you may want to use specific sources for C++ buffers. To do that, register a hook by add-hook and change ac-sources properly:
```
(add-hook 'c++-mode (lambda () (add-to-list 'ac-sources 'ac-source-semantic)))
```
# 7.14. Completion with specific source ===
You can start completion with specific source. For example, if you want to complete file name, do M-x ac-complete-filename at point. Or if you want to complete C/C++ member name, do M-x ac-complete-semantic at point. Usually, you may bind them to some key like:
```
;; Complete member name by C-c . for C++ mode.
(add-hook 'c++-mode-hook
          (lambda ()
            (local-set-key (kbd "C-c .") 'ac-complete-semantic)))
;; Complete file name by C-c /
(global-set-key (kbd "C-c /") 'ac-complete-filename)
```
Generally, such commands will be automatically available when sources are defined. Assume that a source named ac-source-foobar is being defined for example, a command called ac-complete-foobar will be also defined automatically. See also builtin sources for available commands.
If you want to use multiple sources for a command, you need to define a command for it like:
```
(defun semantic-and-gtags-complete ()
  (interactive)
  (auto-complete '(ac-source-semantic ac-source-gtags)))
```
auto-complete function can take an alternative of ac-sources.
# 7.15. Show help persistently ===
Use ac-persist-help instead of ac-help, which is bound to M-<f1> and C-M-?.
# 7.16. Show a lastly completed candidate help ===
ac-last-help command shows a lastly completed candidate help in a ac-help (buffer help) form. If you give an argument by C-u or just call ac-last-persist-help, its help buffer will not disappear automatically.
ac-last-quick-help command show a lastly completed candidate help in a ac-quick-help (quick help) form. It is useful if you want to see a function documentation, for example.
You may bind keys to these command like:
```
(define-key ac-mode-map (kbd "C-c h") 'ac-last-quick-help)
(define-key ac-mode-map (kbd "C-c H") 'ac-last-help)
```
# 7.17. Show help beautifully ===
If pos-tip.el is installed, auto-complete-mode uses its native rendering engine for displaying quick help instead of legacy one.
# 8. Configuration ==
Any configuration item will be set in .emacs or with M-x customize-group RET auto-complete RET.
8.1. ac-delay
Delay time to start completion in real number seconds. It is a trade off of responsibility and performance.
8.2. ac-auto-show-menu
Show completion menu automatically if t specified. t means always automatically showing completion menu. nil means never showing completion menu. Real number means delay time in seconds.
8.3. ac-show-menu-immediately-on-auto-complete
Whether or not to show completion menu immediately on auto-complete command. If inline completion has already been showed, this configuration will be ignored.
8.4. ac-expand-on-auto-complete
Whether or not to expand a common part of whole candidates.
8.5. ac-disable-faces
Specify a list of face symbols for disabling auto completion. Auto completion will not be started if a face text property at a point is included in the list.
8.6. ac-stop-flymake-on-completing
Whether or not to stop Flymake on completion.
8.7. ac-use-fuzzy
Whether or not to use fuzzy matching.
8.8. ac-fuzzy-cursor-color
Change cursor color to specified color when fuzzy matching is started. nil means never changed. Available colors can be seen with M-x list-colors-display.
8.9. ac-use-comphist
Whether or not to use candidate suggestion. nil means never using it and get performance better maybe.
8.10. ac-comphist-threshold
Specify a percentage of limiting lower scored candidates. 100% for whole scores.
8.11. ac-comphist-file
Specify a file stores data of candidate suggestion.
8.12. ac-use-quick-help
Whether or not to use quick help.
8.13. ac-quick-help-delay
Delay time to show quick help in real number seconds.
8.14. ac-menu-height
Specify an integer of lines of completion menu.
8.15. ac-quick-help-height
Specify an integer of lines of quick help.
8.16. ac-candidate-limit
Limit a number of candidates. Specifying an integer, the value will be a limit of candidates. nil means no limit.
8.17. ac-modes
Specify major modes as a list of symbols that will be enabled automatically if global-auto-complete-mode is enabled.
8.18. ac-compatible-packages-regexp
Specify a regexp that identifies starting completion or not for that package.
8.19. ac-trigger-commands
Specify commands as a list of symbols that starts completion automatically. self-insert-command is one of default.
8.20. ac-trigger-commands-on-completing
Same as ac-trigger-commands expect this will be used on completing.
8.21. ac-trigger-key
Specify a trigger key.
8.22. ac-auto-start
Specify how completion will be started. t means always starting completion automatically. nil means never started automatically. An integer means completion will not be started until the value is more than a length of the completion target string.
8.23. ac-ignores
Specify a list of strings that stops completion.
8.24. ac-ignore-case
Specify how distinguish case. t means always ignoring case. nil means never ignoring case. smart in symbol means ignoring case only when the completion target string doesn't include upper characters.
8.25. ac-dwim
"Do What I Mean" function. t means:
After selecting candidates, TAB will behave as RET
TAB will behave as RET only on candidate remains
8.26. ac-use-menu-map
Specify a special keymap (ac-menu-map) should be enabled when completion menu is displayed. ac-menu-map will be enabled when it is t and satisfy one of the following conditions:
ac-auto-start and ac-auto-show-menu are not nil, and completion menu is displayed after starting completion
Completion menu is displayed by auto-complete command
Completion menu is displayed by ac-isearch command
8.27. ac-use-overriding-local-map
Use only when operations is not affected. Internally it uses overriding-local-map, which is too powerful to use with keeping orthogonality. So don't use as much as possible.
8.28. ac-completion-face
Face of inline completion.
8.29. ac-candidate-face
Face of completion menu background.
8.30. ac-selection-face
Face of completion menu selection.
8.31. global-auto-complete-mode
Whether or not to use auto-complete-mode globally. It is t in general.
8.32. ac-user-dictionary
Specify a dictionary as a list of string for completion by dictionary.
8.33. ac-user-dictionary-files
Specify a dictionary files as a list of string for completion by dictionary.
8.34. ac-dictionary-directories
Specify a dictionary directories as a list of string for completion by dictionary.
8.35. ac-sources
Specify sources as a list of source. This is a buffer local variable.
8.36. ac-completing-map
Keymap for completion.
8.37. ac-menu-map
Keymap for completion on completion menu. See also ac-use-menu-map.
8.38. ac-mode-map
Keymap for auto-complete-mode enabled buffers.
# 9. æ©å± ==
æ©å± auto-complete-mode å®éä¸æ¯å®ä¹æ°ç Source.
# 9.1. åå ===
Sourceå¤§è´å¦ä¸:
```
(defvar ac-source-mysource1
  '((prop . value)
    ...))
```
å¶å®å°±æ¯äºK,Vå¯¹ï¼ ä¸é¢ä¼è®²å°é½æé£äºK,ä»ä»¬æ¯ç¨æ¥å¹²åçã
# 9.2. candidates ===
```
(defvar ac-source-mysource1
  '((candidates . (list "Foo" "Bar" "Baz"))))
Then add this source to ac-sources and use:
(setq ac-sources '(ac-source-mysource1))
```
It  is successful if you have "Bar" and "Baz" by inserting "B". The example above has an expression (list ...) in candidates property. The expression specified there will not be byte-compiled, so you should not use an expression unless it is too simple, because it has a bad affection on performance. You should use a function instead maybe:
åºäºå¯¹æ§è½çèèï¼æä»¬æ candidates æ¹åä¸ºå½æ°:
```
(defun mysource1-candidates ()
  '("Foo" "Bar" "Baz"))
(defvar ac-source-mysource1
  '((candidates . mysource1-candidates)))
```
The function specified in candidates property will be called without any arguments on every time candidates updated. There is another way: a variable.
# 9.3. init ===
ä½ å¯ä»¥å¨ç¬¬ä¸æ¬¡è¡¥å¨æ¶ä¸æ¬¡æ§åå§ååéå¯¹è±¡ã
You may want to initialize a source at first time to complete. Use init property in these cases. As same as candidates property, specify a function without any parameters or an expression. Here is an example:
(defvar mysource2-cache nil)
(defun mysource2-init ()
  (setq mysource2-cache '("Huge" "Processing" "Is" "Done" "Here")))
(defvar ac-source-mysource2
  '((init . mysource2-init)
    (candidates . mysource2-cache)))
In this example, mysource2-init function does huge processing, and stores the result into mysource2-cache variable. Then specifying the variable in candidates property, this source prevents huge processing on every time update completions. There are possible usage:
Do require
Open buffers first of all
# Cache ===
Caching strategy is important for auto-complete-mode. There are two major ways: init property and cache property that is described in this section. Specifying cache property in source definition, a result of evaluation of candidates property will be cached and reused the result as the result of evaluation of candidates property next time.
Rewrite the example in previous section by using cache property.
(defun mysource2-candidates ()
  '("Huge" "Processing" "Is" "Done" "Here"))
(defvar ac-source-mysource2
  '((candidates . mysource2-candidates)
    (cache)))
There is no performance problem because this source has cache property even if candidates property will do huge processing.
9.4.1. Cache Expiration
It is possible to keep among more wider scope than init property and cache property. It may be useful for remembering all function names which is rarely changed. In these cases, how can we clear cache property not at the expense of performance? This is true time use that functionality.
Use ac-clear-variable-after-save to clear cache every time a buffer saved. Here is an example:
(defvar mysource3-cache nil)
(ac-clear-variable-after-save 'mysource3-cache)
(defun mysource3-candidates ()
  (or mysource3-cache
      (setq mysource3-cache (list (format "Time %s" (current-time-string))))))
(defvar ac-source-mysource3
  '((candidates . mysource3-candidates)))
Add this source to ac-sources and complete with "Time". You may see a time when completion has been started. After that, you also see the same time, because mysource3-candidates returns the cache as much as possible. Then, save the buffer once and complete with "Time" again. In this time, you may find a new time. An essence of this source is to use ac-clear-variable-after-save to manage a variable for cache.
It is also possible to clear cache periodically. Use ac-clear-variable-every-minute to do that. A way to use is same to ac-clear-variable-after-save except its cache will be cleared every minutes. A builtin source ac-source-functions uses this functionality.
9.5. Action
Complete by RET will evaluate a function or an expression specified in action property. A builtin sources ac-source-abbrev and ac-source-yasnippet use this property.
9.6. Omni Completion
Omni Completion is a type of completion which regards of a context of editing. A file name completion which completes with slashed detected and a member name completion in C/C++ with dots detected are omni completions. To make a source support for omni completion, use prefix property. A result of evaluation of prefix property must be a beginning point of completion target string. Retuning nil means the source is disabled within the context.
Consider a source that completes mail addresses only after "To: ". First of all, define a mail address completion source as same as above.
(defvar ac-source-to-mailaddr
  '((candidates . (list "foo1@example.com"
                        "foo2@example.com"
                        "foo3@example.com"))))
(setq ac-sources '(ac-source-to-mailaddr))
Then enable completions only after "To: " by using prefix property. prefix property must be one of:
Regexp
Function
Expression
Specifying a regexp, auto-complete-mode thinks of a point of start of group 1 or group 0 as a beginning point of completion target string by doing re-search-backward[1] with the regexp. If you want to do more complicated, use a function or an expression instead. The beginning point that is evaluated here will be stored into ac-point. In above example, regexp is enough.
^To: (.*)
A reason why capturing group 1 is skipping "To: ". By adding this into the source definition, the source looks like:
(defvar ac-source-to-mailaddr
  '((candidates . (list "foo1@example.com"
                        "foo2@example.com"
                        "foo3@example.com"))
    (prefix . "^To: \(.*\)")))
Add this source to ac-sources and then type "To: ". You will be able to complete mail addresses.
9.7. ac-define-source
You may use an utility macro called ac-define-source which defines a source and a command.
(ac-define-source mysource3
  '((candidates . (list "Foo" "Bar" "Baz"))))
This expression will be expanded like:
(defvar ac-source-mysource3
  '((candidates . (list "Foo" "Bar" "Baz"))))
(defun ac-complete-mysource3 ()
  (interactive)
  (auto-complete '(ac-source-mysource3)))
A source will be defined as usual and in addition a command that completes with the source will be defined. Calling auto-complete without arguments will use ac-sources as default sources and with arguments will use the arguments as default sources. Considering compatibility, it is difficult to answer which you should use defvar and ac-define-source. Builtin sources are defined with ac-define-sources, so you can use them alone by binding some key to these commands such like ac-complete-filename. See also [this tips](#Completionwithspecific_source].
9.8. Source Properties
9.8.1. init
Specify a function or an expression that is evaluated only once when completion is started.
9.8.2. candidates
Specify a function, an expression, or a variable to calculate candidates. Candidates should be a list of string. If cache property is enabled, this property will be ignored twice or later.
9.8.3. prefix
Specify a regexp, a function, or an expression to find a point of completion target string for omni completion. This source will be ignored when nil returned. If a regexp is specified, a start point of group 1 or group 2 will be used as a value.
9.8.4. requires
Specify a required number of characters of completion target string. If nothing is specified, auto-complete-mode uses ac-auto-start instead.
9.8.5. action
Specify a function or an expression that is executed on completion by RET.
9.8.6. limit
Specify a limit of candidates. It overrides ac-candidate-limit partially.
9.8.7. symbol
Specify a symbol of candidate meaning in one character string. The symbol will be any character, but you should follow the rule:
Symbol	Meaning
s	Symbol
f	Function, Method
v	Variable
c	Constant
a	Abbreviation
d	Dictionary
9.8.8. summary
Specify a summary of candidate in string. It should be used for summarizing the candidate in short string.
9.8.9. cache
Use cache.
9.8.10. require
Specify an integer or nil. This source will be ignored when the integer value is lager than a length of completion target string. nil means nothing ignored.
9.8.11. candidate-face
Specify a face of candidate. It overrides ac-candidate-face partially.
9.8.12. selection-face
Specify a face of selection. It overrides ac-selection-face partially.
9.8.13. depends
Specify a list of features (which are required) that the source is depending.
9.8.14. available
Specify a function or an expression that describe the source is available or not.
9.9. Variables
Here is a list of often used variables.
9.9.1. ac-buffer
A buffer where completion started.
9.9.2. ac-point
A start point of completion target string.
9.9.3. ac-prefix
A string of completion target.
9.9.4. ac-limit
A limit of candidates. Its value may be one of ac-candidate-limit and limit property.
9.9.5. ac-candidates
A list of candidates.
10. Trouble Shooting
10.1. Response Latency
To keep much responsibility is very important for auto-complete-mode. However it is well known fact that a performance is a trade off of functionalities. List up options related to the performance.
10.1.1. ac-auto-start
For a larger number, it reduces a cost of generating completion candidates. Or you can remove the cost by setting nil and you can use when you truly need. See not to complete automatically for more details.
10.1.2. ac-delay
For a larger number, it reduces a cost of starting completion.
10.1.3. ac-auto-show-menu
For a larger number, it reduces a displaying cost of completion menu.
10.1.4. ac-use-comphist
Setting ac-use-comphist to nil to disable candidate suggestion, it reduces a cost of suggestion.
10.1.5. ac-candidate-limit
For a property number, it reduces much computation of generating candidates.
10.2. Completion menu is disrupted
There is two major cases.
10.2.1. Column Computation Case
auto-complete-mode tries to reduce a cost of computation of columns to show completion menu correctly by using a optimized function at the expense of accuracy. However, it probably causes a menu to be disrupted. Not to use the optimized function, evaluate the following code:
(setq popup-use-optimized-column-computation nil)
10.2.2. Font Case
There is a problem when render IPA font with Xft in Ubuntu 9.10. Use VL gothic, which renders more suitably. Or disable Xft, then it can render correctly.
We don't good answers now, but you may shot the troubles by changing font size with set-face-font. For instance, completion menu may be disrupted when displaying the menu including Japanese in NTEmacs. In such case, it is worth to try to evaluate the following code to fix it:
(set-face-font 'ac-candidate-face "MS Gothic 11")
(set-face-font 'ac-selection-face "MS Gothic 11")
11. Known Bugs
11.1. Auto completion will not be started in a buffer flyspell-mode enabled
A way of delaying processes of flyspell-mode disables auto completion. You can avoid this problem by M-x ac-flyspell-workaround. You can write the following code into your ~/.emacs.
(ac-flyspell-workaround)
12. Reporting Bugs
Visit Auto Complete Mode Bug Tracking System and create a new ticket.
# CEDET
# C/CPP
åä¸
ä¸åæäºé½ä»åç å¼å§ã
ç¼ç¨çä¸å¤å¹´ä»£ï¼çæ­£çç å·¥æè®¡ç®æºæä»¤ä»¥ç©¿å­çæ¹å¼æå¨çº¸å¸¦ä¸ï¼ç¶ååç»æºå¨ï¼ç­çå®ååºç»æãåæ¥ï¼IBMåæäºé«çº§ä¸ç¹çæ¹æ³ï¼ç¨ç±»ä¼¼æå­æºçæºæ¢°ï¼keypunch)æäººåæºå¨æå¾çä»£ç åæ¶æå¨ä¸å¼ å¼ å¡çä¸ï¼ä¸å¨å¡çéåèµ·æ¥æ¾å¨çå­éä¹å°±æäºä¸ä¸ªç¨åºãä¿®æ¹ç¨åºæ¶å°±ç´æ¥ç§»å¨æ¿æ¢å¡çãè¿ä¸ªæ¶åï¼ç¼ç¨å¤äºæä½ç©çä»è´¨çé¶æ®µï¼ä¸ä»æ¯èåæ´»ï¼é£çå®ä¹æ¯ä½åæ´»ã
å­åå¹´ä»£ï¼PDPç³»åçæºå¨å¼å§æµè¡äºä¸ï¼ä»PDP-1å¼å§ä¾¿æä¸ä¸ªç¼è¾å¨TECOï¼Tape Editor and COrrector)ãåæçTECOå¶å®æ´åæ¯ä¸ä¸ªçº¯ç²¹çè¯­æ³è§£éå¨ï¼ç¨æ·å¹¶ä¸è½ç¨å®ç´æ¥ç¼è¾ææä»£ç ççº¸å¸¦ï¼èæ¯è¦åæä¿®æ¹çè¿ç¨ç¨å®ç¹æçä¸å¥è¯­è¨åæå¦å¤ä¸ä¸ªä¿®æ­£çº¸å¸¦ï¼åååå§çº¸å¸¦ä¸èµ·åç»TECOï¼å®è§£æä¿®æ­£çº¸å¸¦ä¸çå½ä»¤ç¶åä¾æ¬¡æ§è¡å¨åå§çº¸å¸¦ä¸ï¼æåä¾¿ååºæ¥ä¸ä¸ªæ­£ç¡®ççº¸å¸¦ãTECOæä¾çè¿å¥è¯­è¨ï¼æèè¯´å½ä»¤ï¼ï¼ä»ç¼è¾çè§åº¦è®¾è®¡ï¼çèµ·æ¥æç¹è¯¡å¼ï¼ä½å´å¾å®ç¨ï¼ç±»ä¼¼ bè¡¨ç¤ºå¼å§ï¼cå åæ°è¡¨ç¤ºç§»å¨ï¼iè¡¨ç¤ºæå¥ç­ç­ä¹ç±»ï¼ç¼è¾å¨è¿éä¹å°±æäºæ¹éæ§è¡ç±»ä¼¼æpatchçè¿ç¨ãè¿ä¸ªæ¶åï¼ç¼ç¨è¿ä¸ªå¨ä½æ¬èº«ä¹æ¯ç¼ç¨ï¼è¿æ¯å¤ä¹çº¯çå¹´ä»£ï¼
å­åå¹´ä»£ä¸­æï¼ä»£ç çå­å¨ä»è´¨éæ¸è±ç¦»äºå¡å¸¦ï¼TECOä¹éæ¸åæText Editor and COrrectorã éçPDP-6åè¡çTECO-6ï¼å¯ä»¥å°æä»¶çåå®¹ç´æ¥å±ç¤ºå¨ç»ç«¯ä¸ï¼è¯­è¨çè¯­æ³ä¹åæäºä¸ä¸ªä¸ªçæ­£çå½ä»¤ãç¼è¾æä»¶æ¶ï¼è¾å¥ä¸ä¸ªå½ä»¤ï¼ç¶åæä¸¤ä¸ESCï¼TECOå°å½ä»¤è§£ææ§è¡ï¼å·æ°å±å¹éæ°æ¾ç¤ºåå®¹ï¼ä¿®æ¹çç»æå°±å®æ¶å±ç°äºãè¿æ¶ï¼TECOç»äºæäºç¹WYSIWYG(What You See Is What You Get)çæ¨¡æ ·ï¼ç¨åºç¿ä¹æ»ç®å¯ä»¥å¨å±å¹ååä¸æ¥ï¼å®éçæ²æ²ææäºã
è½¬ä¸
å¦å¼å¼ï¼TECOèµ°è¿äºå­é¶å¹´ä»£ï¼è¿å¥äºä¸é¶å¹´ä»£ãå¨MITçAI Labéï¼ç¨çæºå¨å·²ç»æ¯PDP-6ï¼PDP-10ï¼æºå¨ä¸è·çæ¯ITS(Incompatible Timesharing Systemï¼å«è¿ä¸ªåå­å®å¨æ¯ä¸ªhackä¼ ç»ï¼åªå ä¸ºå®åé¢æä¸ªCompatible Timesharing System), ä¸ç¾¤å¤©æhackerï¼å¤©é«æ°ç½æ¶ä¾¿åç æ¶é£æ¶é£ãè¿ä¸ªæ¶åç¨çTECOå¶å®æ´æç¹åViçæ ·å­ï¼æä½åä¸ºå±ç¤º(display)åç¼è¾(edit)ä¸¤ç§æ¨¡å¼ï¼æ¯å¦ï¼5lå°±æ¯åä¸ç§»å¨5è¡ï¼æä¸iä¹åç´å°ESCä¹é´çè¾å¥è®¤ä¸ºæ¯æå¥çåå®¹ï¼è¾å¥ä¸è¡å­ç¬¦ç¶åæä¸¤ä¸ESCï¼è¾å¥å°±è¢«å½ä½å½ä»¤ä¾æ¬¡æ§è¡ï¼ç¸å½äºViåå·ä¹åæ§è¡å½ä»¤ï¼ãè¿å¶å®æ¯æä¸äºå­ç¬¦çç¹æ®å«ä¹åESCç»åèµ·æ¥ï¼æ ¹æ®å½åçæ¨¡å¼ä¸ºè¾å¥æ¾å°å¯¹åºçæ å°å¤çãå½ç¶ï¼è¿ç§ä¸åçå¤çæ¹å¼è®©æäºhackerè§å¾ä¸ç½ï¼å¶ä¸­çä¸ä½ï¼Carl Mikkelsenï¼ä¾¿ç»ä¸displayåeditæ¨¡å¼ï¼ä¸ºTECOå¢å äºæ´åä¹åçreal-timeæ¨¡å¼ï¼ä¹å«control-Ræ¨¡å¼ï¼ï¼åæ¶ï¼Rici Liknaitskiä¹åæ­¥å®ç°äºæä¸å¨å½ä»¤ç»è£æå¨ä¸èµ·ï¼ç¶åç»å®å°ä¸ä¸ªé®ä¸ãè¿æ¶ï¼control-Rè¿å¥real-timeæ¨¡å¼ä¹åï¼è¾å¥é½è§£ææå½ä»¤ãæ®éå­ç¬¦è¡¨ç¤ºè¾å¥å®æ¬èº«ï¼self-insertingï¼ï¼control metaè¿äºæ§å¶é®åè¡¨ç¤ºç§»å¨ä¿®æ¹ä¹ç±»çæä½ï¼è¿ä¹æ¯ä»ä»¬ç¬¬ä¸æ¬¡è¢«ç¨æ¥ç¼è¾æä»¶ï¼ãå½real-timeæ¨¡å¼æµè¡ä¹åï¼äººä»¬èªç¶æç¹ä¹ä¸æèäºãè¿ä¸ªæ¶åï¼æä»¬çèå¤§Richard Stallmanç»äºåºç°äºï¼ä»æ¥å°AI Labåï¼ä¼åäºreal-timeæ¨¡å¼å¹¶æå®åæTECOçåç½®æ¨¡å¼ï¼å¨Rici Liknaitskiçåºç¡ä¸ä¸ºTECOå¢å äºç¨æ·å¯ä»¥å®å¨èªå®ä¹çå®ï¼ä¹å°±æ¯ä¸å¨å½ä»¤ï¼åè½ï¼åæ¶å¯ä»¥ç»å®å°ä»»æé®ä¸ãæ­¤æ¶ï¼å¤©ä¸ç»å¾ä¸ç»ï¼å¨TECOçç¼éï¼ä¸åè¾å¥çæ¯å®ã
è¿ä¸ªçæ¬çTECOå¤§åæ¬¢è¿ï¼è½èªå®ä¹æ´æ¯æ­£ä¸­é£ä¸ç¾¤hackerçä¸æï¼ä»ä»¬æ ¹æ®èªå·±çéè¦ååå¥½å®ç°ä¸äºèªå®ä¹çå®ï¼åå­é½ä»¥âMACâæèâMACSâç»å°¾æ¥æ è¯ãå¯é®é¢ä¹æ¥äºï¼æ¯ä¸ªäººé½æèªå·±çä¸å¥å®å®ç°æé®ç»å®ï¼ä¸æ¶åå¤©ä¸å¤§ä¹±ï¼ä¸ä¸ªäººåºæ¬ç¨ä¸äºå¦ä¸ä¸ªäººçTECOï¼å ä¸ºæ ¹æ¬å°±ä¸ç¥éè¯¥æä»ä¹é®ãå½æ¶èº«ä¸ºLispæºå¨ï¼Lisp  machineï¼ç»´æ¤åçGuy Steeleï¼è¿æ¯ä¸ªçæ­£éæå¤å½è¯­è¨çç¨åºç¿ï¼FPççåèï¼å ä¸ºè¦ç»å¸¸å¸®äººä»¬å¤çé®é¢æ´æ¯æ·±åå¶å®³ãäºæ¯ï¼ä»åå¤è¦æ¥æ¶æ¾è¿ä¸ªæ··ä¹±å±é¢ãSteeleæ¸¸è¯´åæ¹ï¼è¯´æäººä»¬åæåå¹¶ç»ä¸æ©å±çåç§çæ¬ãå¶å®å¾é¾ä¸è´çæ¯æä¸ªåè½è¦é»è®¤ç»å®å°åªä¸ªé®ä¸ï¼å¹¸äºä»æ¯ä¸ªç¨åºç¿ä¸æ¯æå­åï¼éæ©äºæè®°å¿æç¤ºä¿¡æ¯çç»å®ååï¼è®©é®ä½å­å¨çæçç±ãè¿æ ·ï¼æä»¬å¨è¦è½¬æ¢å¤§å°åæ¶æè½ç«é©¬æ³å°Meta-lï¼lowercaseï¼åMeta-uï¼uppercaseï¼ãSteeleåStallmanèå¤§ä¸èµ·è§èç»ä¸äºåç§å®ççæ¬ï¼å®ä¹å¥½ä¸å¥é»è®¤çé®ä½ç»å®ï¼æ·»å ææ¡£è¯´æï¼ç¶åï¼ä»ä»¬å³å®èªç«é¨æ·ï¼ä¸ä¸ªå«EMACSçç¼è¾å¨è¯çäºã
å³äºEMACSè¿ä¸ªåå­ï¼ç®åç²æ´çèªç¶æ¯Editing MACroSãStallmanèå¤§ä¹è¯´è¿MACSåé¢ç¨Eæ¯å ä¸ºä»æ³è®©EMACSå¯ä»¥æä¸ä¸ªå­ç¬¦çç®ç§°ï¼èEæ¯å½æ¶ITSç³»ç»ä¸è¿æ²¡æè¢«ç¨å°çãå½ç¶ï¼ä¹å¯ä»¥åAI Labéè¿çå°æ·æ·åºEmack & Bolioâsæç¹æ§æ§ï¼å¶ä»çè¯¸å¦Escape-Meta-Alt-Control-Shiftè¿æ ·çè¯¡å¼çæ¬ï¼è¿éä¹æä¸å¨ãEmacsè¯çä¹åï¼ä¹å°±æ¾ç°äºå®å»¶ç»­è³ä»çè®¾è®¡ååï¼The Extensible, Customizable, Self-Documenting Display Editorãåºå±åè½çç»åè®¾è®¡ä¸ºä¸å±æ©å±æä¾äºæå¤§ä¾¿å©åå¯è½ï¼å°å°å­ç¬¦ç§»å¨ï¼å¤§å°ææ¡£æ¸²æï¼Emacsé½æä¾äºä¸°å¯çæ©å±æ¥å£ï¼éç½®ççµæ´»æ§è®©æ©å±èªç±ä½ä¸è³äºå¤±æ§ï¼å¶å®æ©å±æ¬èº«å°±æ¯ä¸ç§éç½®ãè¿ä¸¤ä¸ªçµæ´»æ§ä¹è®©Emacsä¸åä»ä»æ¯ä¸ä¸ªææ¬ç¼è¾å¨ï¼text editorï¼ï¼èæ¯åè½å¼ºå¤§çææ¡£å¤çå¨ï¼word processerï¼ï¼ä¸ºäºè®©æ©å±å¾ä»¥é¦ç«ç»§ä¼ ï¼èªæè¿°çé£æ ¼è®©ä»£ç å³æ¯ææ¡£ãEmacséçææåéï¼å½ä»¤ï¼é®ä½ç¸å³çå¸®å©ææ¡£é½æ¯å¯ä»¥ç´æ¥å¾å°çï¼èè¿äºå¶å®é½æ¯ç´æ¥ä»ä»£ç éæ½ååºæ¥çï¼è¿è·Knuthèå¤§çLiterate programmingåæ¯æç¸éä¹å¤ï¼ã
å½ä¸
èµ·åï¼EMACSçåºå±æ¯ç¨PDP-10çæ±ç¼è¯­è¨å®ç°çï¼ç¨æ·ç¼åèªå®ä¹å®æä½¿ç¨çä»ç¶æ¯TECOæå¼å§çé£å¥è¯­è¨ãä¸ºäºæä¾æ´å¤§ççµæ´»æ§ï¼Statllmanèå¤§ä¸ºè¿å¥å·ç§°ä¸çä¸æåæçè¯­è¨æ·»å äºè®¸å¤å¤æåè½ãä½è¿ä¸ªè¯­è¨ä»æå¼å§å°±åªæ¯é¢åç¼è¾å¨è®¾è®¡ï¼èä¸æ¯é¢åç¨åºè®¾è®¡çãæ·»å çåè½è¶å¤ï¼å®å°±æ¾çè¶è¯¡å¼æ æ¯ãå½è¿ä¸ªâas ugly as possibleâçè¯­è¨è®©èå¤§å´©æºæçæ¶åï¼ä»è®¤è¯å°ç»§ç»­ä½¿ç¨å®å°æ¯ä¸æ¡éè¯¯çéè·¯ï¼å¿é¡»æ¾å°ä¸ä¸ªçæ­£çç¨åºè®¾è®¡è¯­è¨ä½ä¸ºEMACSæ©å±çå¼åéæ©ãè¿ä¸ªæ¶åï¼å·²ç»æäºä¸äºåç§çå®ç°ãDan Weinreb å¨Lispæºå¨ä¸å®ç°ä¸ä¸ªEINEï¼EINE is Not Emacs), Bernard Greenbergä¹å®ç°äºMacLispççæ¬ãè¿ä¸¤ä¸ªçæ¬é½æ¯å®å¨ä½¿ç¨Lispè¯­è¨å®ç°åºå±åå®çå½ä»¤æ©å±ï¼è¿ç»äºStallmanèå¤§éæ°è®¾è®¡EMACSä¸äºå¯åãLispå½æ¶å·²ç»æäºgcæºå¶ï¼éç¨å®ç°å·²ç»ä¸æé®é¢ï¼ä½é¤äºå¨ä¸ç¨çLispæºå¨ä¸å¤ï¼æ§è½ä»ç¶æ¯Lispæè±ä¸äºçè½¯èï¼èå½æ¶cå¨Unixä¸çä¼å¼è¡¨ç°å¸å¼äºèå¤§çæ³¨æãäºæ¯ï¼æè¡¡ä¸æï¼Stallmanèå¤§éåäºæä¸­ç­ç¥ï¼å®ä¸äºæ°çEMACSçè®¾è®¡ï¼cå®ç°åºå±ï¼Lispå®ç°æ©å±ãæ­¤æ¶å·²ç»æ¯å«åå¹´ä»£åï¼MIT AI Labéçhackçæ¬¢æ©å·²ç»æï¼ä¸ºäºèªç±çæ³ï¼åæªå¹é©¬ä¸åä¸çSymbolicsè¦æäºä¸¤å¹´ä¹åï¼Stallmanèå¤§ä¹è®¤è¯å°è¡¥æ§ä¸å¦ç«æ°ï¼çæå¼å§ä»åååä»£çGNUé¡¹ç®ï¼èæ°ççEMACSæäºGNUé¡¹ç®çåè½«ä¹ä½ã
å¶å®ï¼James Goslingäº1981å¹´å·²ç»å¨Unixä¸å®ç°äºGosling Emacsï¼å®çåºå±èªç¶æ¯Unixä¸å¤©ççcè¯­è¨ï¼å½ä»¤æ©å±ä½¿ç¨çæ¯ä¸ä¸ªè¯­æ³ç±»LispçMockLispå®ç°ãStallmanèå¤§åèäºä¸äºGosling Emacså®ç°ï¼ä½ä¸çº¯æ­£çLispå®ç°è¿æ¯ç¸å½ä¸å¥èå¤§çæ³ç¼ï¼åæ­£å½æ¶çLispä¹æ²¡ææ åè§èï¼ä¸å¨å¨çæ¹è¨æ¼«å¤©é£ãèå¤§æç§ä¸¥æ ¼çLispè¯­ä¹ï¼å®å¨éåäºä¸ä¸ªfull-featureçLispçæ¬ï¼è¿ä¹å°±æ¯æ²¿ç¨è³ä»çElispï¼Emacs Lispï¼ãèµ·åGosling Emacsä¹æ¯ä½ä¸ºèªç±è½¯ä»¶èªç±åè´¹å°ä¼ æ­ï¼å½æ¶Stallmanèå¤§ä¹æ­£å¨æ¨è¡Emacsåè°ï¼Emacs commune)è¿å¨ï¼å®£æ¬ä»çèªç±è½¯ä»¶çæçæ³ï¼äºèä¹é´å¼å§è¿æ¯å¾æé»å¥çãä½éçGosling Emacsä½¿ç¨èå´çæ©å¤§ï¼1984å¹´ï¼Goslingå£°æèªå·±å·²ç»æ²¡æè½åç»§ç»­ç»´æ¤ä¸å»ï¼å°å¶åç»äºå½æ¶çåä¸è½¯ä»¶å¬å¸Unipressãè¿ä¸ªè¡ä¸ºæ¬æ¯æ å¯åéï¼æ éæ¯ä»£ç çæçå½å±ååè¡æ¹å¼çæ¹åï¼ä½å¨Stallmanèå¤§ç¼éï¼è½¯ä»¶ä»£ç æ¯èªç±çæ³çè½½ä½ï¼æµä¼ åè¡çæ¹å¼æä¸å±äººç±»ç¤¾ä¼çå­çå²å­¦å«ä¹ãGoslingæåç¸æ¯è¿å»çè¡ä¸ºæ å¼äºåå¾ç¹å¤§ï¼å®å¨æ¯å¯¹ä¿¡ä»°çäºµæ¸ãèå¤§ä¸¥è¾èè§Goslingä¸ºæ¦å¤«ï¼å¿å°éå¨åå²è»è¾±æ±ä¸ä¹ç±»ç§ç§ãä½èå¤§çè¡ä¸ºæ»æ¯æ¯æ²å£®çï¼Gosling Emacsè¿æ¯æä¸ºäºä¸ä¸ªåä¸è½¯ä»¶ï¼James Goslingåæ¥ä¹å»äºSunï¼æå°±äºä»ä¼ ä¸çJavaä¼ä¸ã
1985å¹´ï¼æ°çæ¬çEmacsç»äºä½ä¸ºGNU Emacsç¬¬ä¸æ¬¡å¬å¼åè¡ï¼çæ¬å·å®ä¸º13ãè³äºçæ¬1å°12æ¯ä»æ¥æ²¡æåè¡è¿çï¼å ä¸ºå½æ¶å¼åçæ¬ä¸ç´ä½¿ç¨1.xxçæ ¼å¼, ä½å¨1.12ä¹åï¼ä»è®¤ä¸ºGNU Emacsä¸»çæ¬å·1æ¯æ°¸è¿ä¸ä¼åçï¼éä¹ä¾¿æå®å»æäºãåæ¥çåè¡çæ¬ä¹å°±ä»13å¼å§å»¶ç»­ä¸æ¥ï¼ä¸ç´å°ç°å¨ãVesion 15çGNU Emacså·²ç»å¯ä»¥å®ç¾çå¨Unixç³»åä¸è¿è¡ï¼å®å®å¨çLispåè½å®ç°ä¸èªç±åè´¹çä¼ æ­ä½¿ç¨ï¼èªç¶æ¯åä¸çGoling Emacsï¼Unipress Emacsï¼æ´æå¸å¼åãä½è¿æ¶éº»ç¦ä¹æ¥äºï¼GNU Emacså¨æ¾ç¤ºé¨åä½¿ç¨çä¸äºGosling Emacsä»£ç ï¼ç°å¨å¼æ¥äºåä¸çæä¸é¢çå²çªãè¿è®©Stallmanèå¤§å¾æ¼ç«ï¼ä»é¦åå£°æï¼ä½¿ç¨çä»£ç æ¯éè¿Fen Labalmeï¼å½æ¶åJames Golingä¸èµ·å¼åGosling Emacsï¼æGoslingæç»çä»£ç ä½¿ç¨æï¼ï¼å®å¨æ¯åæ³çãåæ¥ï¼äºæå½ç¶ä¸ä¼å¦ç°å¨æ­¤ç±»äºæé£æ ·æ è¾¹ççº ç¼ ä¸å»ï¼å¤çç±»ä¼¼æ¯ç®é®é¢ï¼èä¸è¾ä»¬æçä¸è´¯ç®åç´æ¥çå¤çæ¹æ³ï¼ä½ è¯´ä»£ç æ¯ä½ çï¼æä¸ç¨å°±å¥½äºãStallmanèå¤§ç´æ¥å£°æâI have decided to replace the Gosling code in GNU Emacs, even though I still believe Fen and I have permission to distribute that code ãããI expect to have the job done by the weekendãããâãä¸ä¸ªææä¹åï¼èå¤§éåäºæäºè®®çä»£ç å®ç°ï¼Version 16ä»¥åçGNU Emacsæäºçæ­£çGosling-freeçæ¬ãä½è¿ä»¶äºè®©Stallmanèå¤§ç¡®å®å¾åä¼¤ï¼è®¤è¯å°GNUè½¯ä»¶å¿é¡»è¦æä¸ä¸ªèªå·±ççææ¥ä¿æ¤ãåæ¥ä¸èµ°å¯»å¸¸è·¯çèå¤§ä¹å°±æ³å°äºæµéhackæ°å³çâcopyleftâï¼GNU Emacsçâthe GNU Emacs copying permission noticeâä¹æäºå®ç¬¬ä¸æ¬¡å®éªæ§çå®ç°ãåå±å°åæ¥ä¸æ®µæ¶é´ï¼æ¯ä¸ªGNUçè½¯ä»¶é½æèªå·±çä¸ä¸ªçæï¼ç±»ä¼¼Emacs General Public Licenseï¼Nethack General Public Licenseã1989å¹´1æï¼ç»äºæ¼ååºäºç¬¬ä¸ä¸ªç»ä¸ççæï¼the General Pulic License version 1ï¼GPL V1ï¼ã
1991å¹´ï¼GNU Emacså·²ç»å°äºversion 19ãå½æ¶å¨Lucidå¬å¸ï¼ä¸å¸®å®¶ä¼è¦æGNU Emacsæ´åæä¸ä¸ªC++çIDEï¼ä¸ºäºæ»¡è¶³éè¦ï¼ä»ä»¬åäºå¾åä¿®æ¹ï¼æ·»å äºè®¸å¤æ°åè½ï¼çè³è¶è¿äºGNU Emacså®æ¹ççæ¬æ´æ°éåº¦ãæåï¼ä»ä»¬ä¸åå¤è·éå®æ¹çèæ­¥äºï¼äºæ¯ä¸ä¸ªæ°çEmacs åç§Lucid Emacsåºç°äºãå½ç¶ï¼Lucidè¿ä¸ªå¾å®¹æå¼èµ·çæé®é¢çåå­åæ¥åæäºçç¥çXEmacsï¼ä¸è·¯åå±ï¼æäºé¤GNU Emacså¤ææµè¡çEmacsçæ¬ãå¦å¤çä¸»æµæä½ç³»ç»å¹³å°ä¸ä¹ä»GNU Emacsä¸è¡ååºäºç¸åºççæ¬ï¼MS WindowsçMeadow, Apple Macintoshä¸çAquamacsã
åæ¯ä¸GNU Emacsçå¯å¨ç»é¢ï¼è¿æ¯å¨2000å¹´çVersion21ä¸­å å¥çãèä¸è¾ä»¬è¿æ¯å¾æèºçï¼Emacsåé¢ççåå­å¶å®æ¯âGNUâçé£æ¸ºåæ³ï¼æåæ¯Gnusï¼å®ç¾ç»åå¨Emacséï¼çlogoãåé¢çâEmacsâæåçæ¬è¦èå¹»å¾å¤ï¼åªæ¯æ²¡è½è®¨å¾Stallmanèå¤§æ¬¢å¿ï¼ææ¹æäºç°å¨çæ ·å­ï¼è¿éælogoçè¯¦ç»åå²æº¯æºï¼ã
ç°ä¸
æ¶å°å¦ä»ï¼å¨Emacs WikiåEmacs Lisp Listä¸å·²ç»æäºæåä¸ä¸çåæä¸åæçæ©å±ï¼æ»¡è¶³åç§åæä¸åæçéè¦ï¼è¿éæä¸ªIs there anything Emacs can not doççæ¯è®¨è®ºï¼ãâEmacs is my operating system, and Linux is its device driverâçå¤è¯­ä¹ä¸æ¯æ²¡æåå çï¼Viä¸æç»å¸¸è¯´çâEmacs is a nice operating system, but lacks a good editorâä¹æ¯å¯ä»¥çè§£çãä½é«éç½®æ§è®©Emacsè½ç¶åºå¤§ä½å¹¶ä¸èè¿ï¼è´¤æ çå¥¹æ¯ï¼ä½ ä¸è¦çï¼å¥¹ä¸ä¼ç»ä½ ï¼ä½ è¦çï¼å¥¹è¯å®æ¯å¯ä»¥ç»ä½ çãEmacsåViæä¸º*nixç¨åºç¿ççåºåæ å¿ï¼è½ç¶Emacséä¹æä¸ªvi modeçï¼ï¼ä½æ¯ï¼Stallmanèå¤§è®¤ä¸ºè¿è¿æ¯ä¸å¤çï¼ä»èªå·±æ¯ä»æ¥æ²¡æå»ç¨è¿Viç)ï¼ä»æé£äºä»æªç¨è¿Emacsçäººç§°ä¸ºEmacs Virginsï¼èå¸¦é¢è¿äºäººèµ°åºè¿ç§å°´å°¬å¢å°æ¯Emacsersçç¥å£ä½¿å½ï¼blessed actï¼ãèå¤§è¿çªè¯å¾æååçï¼æè¾è¿éåªåäºã
ä¸åæäºè¿å¨åç ä¸­ç»§ç»­ãå½ç¶ï¼ä¸åä¹ä¸åªå¨åç ä¸­ç»§ç»­ããã
[[TOC]]
# Haskell
# Haskell Mode
```sh
$ git clone https://github.com/amas-git/haskell-mode
```
# ghc-mod
ä»cabalå®è£
```
$ cabal update
$ cabal install ghc-mod structured-haskell-mode stylish-haskell
```
å®è£å®æ¯å, ä¼å¾å°ä»¥ä¸å ä¸ªæä»¶
 * ~/.cabal/bin/ghc-mod ghc-modi
 * ~/.cabal/share/i386-linux-ghc-7.8.3/ghc-mod-4.1.6/ : Emacsæ¯æ
```
$ cd ~/.emacs.d
$ ln -s ~/.cabal/share/i386-linux-ghc-7.8.3/ghc-mod-4.1.6/  ghc-mod
```
å¨ä½ ç.emacsæä»¶ä¸­å å¥åå§åä»£ç :
```el
(add-to-list 'load-path "~/.emacs.d/ghc-mod/")
(autoload 'ghc-init "ghc" nil t)
(autoload 'ghc-debug "ghc" nil t)
(add-hook 'haskell-mode-hook (lambda () (ghc-init)))
```
åºæ¬æä½: http://www.mew.org/~kazu/proj/ghc-mod/en/emacs.html
# auto-complete
æä½æ¥æ¬æååäº«äºæ¯è¾å®åçauto-complteçhaskelléç½®: http://www.mew.org/~kazu/proj/ghc-mod/en/emacs.html
```
$ ghc-mod
ghc-mod version 4.1.6 compiled by GHC 7.8.3
Usage:
         ghc-mod list [-g GHC_opt1 -g GHC_opt2 ...] [-l] [-d]
         ghc-mod lang [-l]
         ghc-mod flag [-l]
         ghc-mod browse [-g GHC_opt1 -g GHC_opt2 ...] [-l] [-o] [-d] [-q] [<package>:]<module> [[<package>:]<module> ...]
         ghc-mod check [-g GHC_opt1 -g GHC_opt2 ...] <HaskellFiles...>
         ghc-mod expand [-g GHC_opt1 -g GHC_opt2 ...] <HaskellFiles...>
         ghc-mod debug [-g GHC_opt1 -g GHC_opt2 ...] 
         ghc-mod info [-g GHC_opt1 -g GHC_opt2 ...] <HaskellFile> <module> <expression>
         ghc-mod type [-g GHC_opt1 -g GHC_opt2 ...] <HaskellFile> <module> <line-no> <column-no>
         ghc-mod find <symbol>
         ghc-mod lint [-h opt] <HaskellFile>
         ghc-mod root
         ghc-mod doc <module>
         ghc-mod boot
         ghc-mod version
         ghc-mod help
<module> for "info" and "type" is not used, anything is OK.
It is necessary to maintain backward compatibility.
  -l           --tolisp             print as a list of Lisp
  -h hlintOpt  --hlintOpt=hlintOpt  hlint options
  -g ghcOpt    --ghcOpt=ghcOpt      GHC options
  -o           --operators          print operators, too
  -d           --detailed           print detailed info
  -q           --qualified          show qualified names
  -b sep       --boundary=sep       specify line separator (default is Nul string)
```
# åè
 * http://sritchie.github.io/2011/09/25/haskell-in-emacs/
[[TOC]]
# Lazy Evaluation
åå¦æå¦ä¸å½æ°:
```hs
> let fun x y = x
> fun 1 2
1
> fun 1 (1 `div` 0)
1
```
Lazy Evaluation åªå¨éè¦å¯¹è¡¨è¾¾å¼æ±å¼çæ¶åæå»æ±å¼. åä¹EagerEvaluationæ»æ¯å¯¹è¡¨è¾¾å¼è¿è¡æ±å¼.
```div class=note
# é¤é¶éè¯¯
```hs
> (1/0)
Infinity
>1 `div` 0
*** Exception: divide by zero
```
ä¸ºä»ä¹å¢ï¼ ( http://stackoverflow.com/questions/9354016/division-by-zero-in-haskell )
```
```
```
# Org Mode
 * http://orgmode.org/    
# ç§»å¨
|| C-c C-n  || ä¸ä¸ä¸ªæ é¢ 
|| C-c C-p  || åä¸ä¸ªæ é¢
|| C-c C-f  || ä¸ä¸ä¸ªåçº§æ é¢
|| C-c C-b  || åä¸ä¸ªåçº§æ é¢
|| C-c C-u  || åå°ç¶æ é¢
# ç»æç¼è¾
|| M-return    || æå¥åçº§æ é¢
|| M-S-return  || æå¥åçº§TODO
|| M-left      || éä½æ é¢çº§å«
|| M-right     || æé«æ é¢çº§å«
|| M-S-left    || éä½æ é¢çº§å«(åæ¬å­æ é¢)
|| M-S-right   || éä½æ é¢çº§å«(åæ¬å­æ é¢)
|| M-up        || åä¸ç§»å¨æ é¢
|| M-down      || åä¸ç§»å¨æ é¢
|| M-S-up      || åä¸ç§»å¨æ é¢(åæ¬å­æ é¢)
|| M-S-down    || åä¸ç§»å¨æ é¢(åæ¬å­æ é¢)
# Refile(å¨èç¹ä¹é´ç§»å¨æ é¢æ )
|| C-c C-w     || å°æ é¢æ ç§»å¨å°å¶ä»æ é¢ä¸
# Sparseæ 
æ é¢ï¼TODOè¯¸å¤æ¡ç®æ··æå¨ä¸èµ·çæ¶åï¼æä»¬éè¦ä¸ç§å¿«éæç´¢æéä¿¡æ¯çæ¹æ³.
|| C-c /   || éæ©åå»ºSparseTreeçæ¹æ³
|| C-c / r || ä½¿ç¨RegexpçæSparseTree
|| C-c C-c || è¿å 
|| Tab        || æ¶èµ·
|| M-return   || insert new at same level
|| M-S-return || insert new item with checkbox
|| C-c C-c    || toggle checkbox item
|| C-c -      || åæ¢åè¡¨æ ·å¼
# è¿æ¥
```
[[Linkage][Name]]
```
 * [[http://www.baidu.com][ç¾åº¦]]
 
# TODO
|| C-c C-t ||
|| S-right ||
|| S-left  ||
å¯ä»¥èªå®ä¹TODOæ¡ç®çç¶æ:
```el
(setq org-todo-keywords
      '((sequence "TODO" "FEEDBACK" "VERIFY" "|" "DONE" "DELEGATED")))
```
```el
(setq org-todo-keywords
      '((sequence "TODO(t)" "|" "DONE(d)")
        (sequence "REPORT(r)" "BUG(b)" "KNOWNCAUSE(k)" "|" "FIXED(f)")
        (sequence "|" "CANCELED(c)")))
```
[[TOC]]
# PythonçEmacså¼åç¯å¢
# ä¾èµ
 * emacs23+
 * pyflymkes
 * pymacs
#  å®è£pymacs
 * https://github.com/pinard/Pymacs
1. å®è£Pythoné¨å
```sh
# ä½¿ç¨é»è®¤çæ¬çpython
$ [sudo] make install
# ä½¿ç¨æå®çæ¬çpython, æ¯å¦python2
$ sudo make install PYTHON=python2
# æ¥çå¶ä»ç¼è¯éé¡¹
$ python setup.py install --help
```
æ³è¦ç¡®ä¿¡å®è£æå, è¿å¥å°Pythonä¸­`from Pymacs import lisp`, æ å¼å¸¸å³è¡¨æå·²ç»å®è£äºPymacs.
```
Python 2.7.3 (default, Dec 22 2012, 21:27:36) 
[GCC 4.7.2] on linux2
Type "help", "copyright", "credits" or "license" for more information.
>>> from Pymacs import lisp
```
2. å®è£Emacsé¨å
å¨`~/.emacs`ä¸­å å¥:
```
(add-to-list 'load-path (expand-file-name "~/.emacs.d/python/python/Pymacs"))
(require 'pymacs)
(autoload 'pymacs-apply "pymacs")
(autoload 'pymacs-call  "pymacs")
(autoload 'pymacs-eval  "pymacs" nil t)
(autoload 'pymacs-exec  "pymacs" nil t)
(autoload 'pymacs-load  "pymacs" nil t)
(autoload 'pymacs-autoload "pymacs")
```
å¯å¨emacs, ç¨pstreeè§å¯ä¸ä¸, Pythonèç¹èµ·æ¥äº, è¯´æä¸åæ­£å¸¸.
```
...
ââemacs --name main.edit
â   ââpython2 -c import sys; from Pymacs import main; main(*sys.argv[1:]) -f
â   ââ{emacs}
...
```
# å®è£rope
```
$ wget https://pypi.python.org/packages/source/r/rope/rope-0.9.4.tar.gz
$ cd rope-0.9.4
$ [sudo] python2 setup.py install
```
# å®è£ropemacs
 * https://bitbucket.org/agr/ropemacs
`~/.emacs`å å¥:
```
(setq ropemacs-enable-autoimport t)
(pymacs-load "ropemacs" "rope-")
```
# å®è£python-mode
 * https://launchpad.net/python-mode
1. `~/.emacs`å å¥
```el
(add-to-list 'load-path "~/.emacs.d/python/python-mode") 
(setq py-install-directory "~/.emacs.d/python/python-mode")
(require 'python-mode)
(setq py-shell-name "python2")        ; é»è®¤python shell
(setq-default indent-tabs-mode nil)   ; ç¼©è¿æ¶åªä½¿ç¨ç©ºæ ¼
```
å¸¸ç¨çåè½:
||= å¿«æ·é® =||= åè½  =||= è¯´æ =||
|| C-c C-c    ||  è¿è¡å½åæä»¶ || py-execute-buffer || 
# å®è£python.el (`ææºä¼åè¯è¯`)
[https://github.com/fgallina/python.el python.el]æ¯å¦å¤ä¸ä¸ªæ¯è¾å¥½ç¨çä¸»æ¨¡å¼.
å°åå«å¨Emacs24.3ä¹åçåè¡çä¸­. ç®åä¸»çº¿ä»æ¯æEmacs23, Emacs24.2çåå­¦éè¦åæ¢å°å¯¹åºçåæ¯.
# åè
 * http://edward.oconnor.cx/2008/02/ropemacs
 * http://www.saltycrane.com/blog/2010/05/my-emacs-python-environment/
# Yasnippet =
 * http://code.google.com/p/yasnippet/
# æºç 
```sh
$ git clone https://github.com/capitaomorte/yasnippet
```
# éç½® 
```el
;;------------------------------------------------------------------[ Yasnippet ]
(add-to-list 'load-path (expand-file-name "~/.emacs.d/yasnippet"))
(require 'yasnippet)
(yas/initialize)
(yas/load-directory "~/.emacs.d/yasnippet/snippets")
```
# æ©å±æ¨¡æ¿
ä½ å¯ä»¥åå°snippet-modeç¼è¾æ¨¡æ¿æä»¶.
# M-x yas/load-snippet-buffer
 * å½ä½ ç¼è¾æ¨¡æ¿æ¶,å¯ä»¥ä½¿ç¨æ­¤å½ä»¤å°æ¨¡æ¿è£å¥èåä¸­
 * é»è®¤ç»å®å°C-c C-c
# M-x yas/tryout-snippet
 * æå¼ä¸ä¸ªæ°çbuffer, å¶ä¸­æå¥æ¨¡æ¿å®ä¾, ä½ å¯ä»¥èæ­¤ççä½¿ç¨ææ. 
 * é»è®¤ç»å®å°C-c C-t
# æ¨¡æ¿å¤´
æ¨¡æ¿å¤´å¤§æ¦æ¯ä¸é¢è¿ä¸ªæ ·å­
```
#name:
# ... 
# ...
# -- 
æ¨¡æ¿åå®¹
```
Here's a list of currently supported directives:
# # key: snippet abbrev
This is the probably the most important directive, it's the abbreviation you type to expand a snippet just before hitting yas/trigger-key. If you don't specify this the snippet will not be expandable through the key mechanism.
# # name: snippet name
This is a one-line description of the snippet. It will be displayed in the menu. It's a good idea to select a descriptive name for a snippet -- especially distinguishable among similar snippets.
If you omit this name it will default to the file name the snippet was loaded from.
# # condition: snippet 
#contition: ä¹åå¯ä»¥æå®ä¸æ®µELispä»£ç , æ¨¡æ¿ä»å¨ELispä»£ç è¿åéç©ºæ¶è¿è¡æ©å±.
This is a piece of Emacs-lisp code. If a snippet has a condition, then it will only be expanded when the condition code evaluate to some non-nil value.
See also yas/buffer-local-condition in Expanding snippets
# # group: snippet menu grouping
When expanding/visiting snippets from the menu-bar menu, snippets for a given mode can be grouped into sub-menus . This is useful if one has too many snippets for a mode which will make the menu too long.
The # group: property only affect menu construction (See the YASnippet menu) and the same effect can be achieved by grouping snippets into sub-directories and using the .yas-make-groups special file (for this see Organizing Snippets
Refer to the bundled snippets for ruby-mode for examples on the # group: directive. Group can also be nested, e.g. control structure.loops tells that the snippet is under the loops group which is under the control structure group.
# # expand-env: expand environment
ä¸æ®µElispä»£ç , å½æ¨¡æ¿æ©å±ä¹å,è¯¥æ®µä»£ç ä¼è¢«æ±å¼, æä»¥ä½ å¯ä»¥èæ­¤æ¹åä¸äºéç½®åæ°.
This is another piece of Emacs-lisp code in the form of a let varlist form, i.e. a list of lists assigning values to variables. It can be used to override variable values while the snippet is being expanded.
Interesting variables to override are yas/wrap-around-region and yas/indent-line (see Expanding Snippets).
As an example, you might normally have yas/indent-line set to 'auto and yas/wrap-around-region set to t, but for this particularly brilliant piece of ASCII art these values would mess up your hard work. You can then use:
# # binding: direct keybinding
ç»å®æ¨¡æ¿å°é®åºä¸.
```
#name : <p>...</p>
#binding: C-c C-c C-m
# --
<p>`(when yas/prefix "
")`$0`(when yas/prefix "
")`</p>
```
# # --
æ è®°æ¨¡æ¿åå®¹å¼å§ä¹å¤
# æ¨¡æ¿è¯­è¨
# $
# `
# \
# åµå¥ESlispä»£ç : `(<elisp-code>)`
Cå¤´æä»¶æ¨¡æ¿å¯ä»¥è¿æ ·å®ä¹:
```
#ifndef ${1:_`(upcase (file-name-nondirectory (file-name-sans-extension (buffer-file-name))))`_H_}
#define $1
$0
#endif /* $1 */
```
å¦æä½ çæ¨¡æ¿ä¸­ä¸é¨ååå®¹ä¾èµäºéä¸­æå­,ä½ å¯ä»¥ä½¿ç¨`yas/selected-text`å½æ°:
```
for ($1;$2;$3) {
  `yas/selected-text`$0
}
```
# Tabæ¸¸æ 
 * ä½ çæ¨¡æ¿æå¾å¤å¾å¡«ååºå,å½ä½ æTabé®æ¶,å°åæ¢å°ä¸ä¸ä¸ªå¡«ååºå,S-Tabåè¿åä¸ä¸ä¸ªå¡«ååºå.
 * $0: æ¨¡æ¿æ©å±å®æ¯ååæ æç»çä½ç½®
# å ä½ç¬¦
```
${N:default-value}
```
 * N: ä»»ææ´æ°, 0æç¹æ®å«ä¹, å³æ ç¤ºæ¨¡æ¿æ©å±å®æ¯, åæ æç»å¤äºçä½ç½®
 * default-value: å ä½ç¬¦çé»è®¤å¡«åæå­
# å ä½ç¬¦éå
å¨æ¨¡æ¿ä¸­ææ$Né½æç¸åçå«ä¹, å½ä½ è¾å¥æ¶ææ$Né½ä¼éä¹äº§çç¸åçåå.
# å ä½ç¬¦éååæ¢
default-valueåºç°`$()`, æ¬å·ä¸­çåå®¹å°è¢«ä½ä¸ºELispè¡¨è¾¾è¯è¿è¡æ±å¼
```
${2:hello}
${2:$(capitalize text)}
```
# ä»åè¡¨ä¸­éæ©å ä½ç¬¦çå¼
```
<div align="${2:$$(yas/choose-value '("right" "center" "left"))}">
  $0
</div>
```
# å ä½ç¬¦åµå¥
```
<div${1: id="${2:some_id}"}>$0</div>
```
This allows you to choose if you want to give this div an id attribute. If you tab forward after expanding it will let you change "some_id" to whatever you like. Alternatively, you can just press C-d (which executes yas/skip-and-clear-or-delete-char) and go straight to the exit marker.
By the way, C-d will only clear the field if you cursor is at the beginning of the field and it hasn't been changed yet. Otherwise, it performs the normal Emacs delete-char command.
# èªå¨å¯¹é½: $>
```
for (${int i = 0}; ${i < 10}; ${++i})
{$>
$0$>
}$>
```
# .yas.parents æä»¶
ä½ å¯ä»¥å¨`.yas-parents`æä»¶ä¸­æå®å¤ç¨ç¶æ¨¡æ¿.
```sh
$ tree
.
|-- c-mode
|   |-- .yas-parents    # contains "cc-mode text-mode"
|   `-- printf
|-- cc-mode
|   |-- for
|   `-- while
|-- java-mode
|   |-- .yas-parents    # contains "cc-mode text-mode"
|   `-- println
`-- text-mode
    |-- email
```
# å¸¸ç¨ELisp
```el
(file-name-nondirectory (file-name-sans-extension (buffer-file-name)))
```
```el
(format-time-string "%Y-%m-%d %H:%M" (current-time))
```

