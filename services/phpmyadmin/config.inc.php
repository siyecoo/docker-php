<?php

# 配置文件从 phpmyadmin 容器中获得:
# docker run --rm phpmyadmin:latest cat /etc/phpmyadmin/config.inc.php >> config.inc.php

# 编辑 config.inc.php
# 查找 $cfg['Servers'][$i]['AllowNoPassword'] 行
# 并添加以下行，ruser 替换自己的只读用户

# $i = 1;
# $cfg['Servers'][$i]['AllowRoot'] = false;
# $cfg['Servers'][$i]['AllowDeny']['order'] = 'explicit';
# $cfg['Servers'][$i]['AllowDeny']['rules'] = ['allow ruser from all'];