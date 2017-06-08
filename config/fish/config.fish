# 环境变量
set -x GOPATH ~/Projects/go
set -x PATH $GOPATH/bin ~/.local/bin $PATH
set -x EDITOR vim
set -x LC_CTYPE zh_CN.UTF-8     # 中文处理
set -x LC_COLLATE zh_CN.UTF-8   # 中文排序

# 提示符
function fish_prompt
  # Just calculate these once, to save a few cycles when displaying the prompt
  if not set -q __fish_prompt_hostname
    set -g __fish_prompt_hostname (prompt_hostname)
  end
  if not set -q __fish_prompt_char
    switch (id -u)
      case 0
	      set -g __fish_prompt_char '#'
      case '*'
	      set -g __fish_prompt_char '>'
    end
  end

  # Setup colors
  set -l normal (set_color normal)
  set -l white (set_color FFFFFF)
  set -l turquoise (set_color 5fdfff)
  set -l orange (set_color df5f00)
  set -l hotpink (set_color df005f)
  set -l blue (set_color blue)
  set -l limegreen (set_color 87ff00)
  set -l purple (set_color af5fff)
  set -l yellow (set_color yellow)

  # Configure __fish_git_prompt
  set -g __fish_git_prompt_char_stateseparator ' '
  set -g __fish_git_prompt_color 5fdfff
  set -g __fish_git_prompt_color_flags df5f00
  set -g __fish_git_prompt_color_prefix white
  set -g __fish_git_prompt_color_suffix white
  set -g __fish_git_prompt_showdirtystate true
  set -g __fish_git_prompt_showuntrackedfiles true
  set -g __fish_git_prompt_showstashstate true
  set -g __fish_git_prompt_show_informative_status true

  echo -n $hotpink$USER$white'@'$yellow$__fish_prompt_hostname$white
  echo -n ' in '$limegreen(prompt_pwd)$turquoise
  __fish_git_prompt " (%s)"
  echo

  echo -n $white(date +'%F %T')' '$__fish_prompt_char' '$normal
end

function fish_right_prompt
  set -l exit_code $status
  if test $exit_code -ne 0
    set_color red
  else
    set_color 666666
  end
  echo -n '['$exit_code']'

  set_color normal
end

# 别名
alias yxapp="ssh kaizhang33@relay.creditease.corp"

# rbenv
status --is-interactive; and source (rbenv init -|psub)
