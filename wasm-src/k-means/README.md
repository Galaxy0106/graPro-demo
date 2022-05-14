## Prerequisite
```shell
# 选择默认安装
curl https://sh.rustup.rs -sSf | sh
# 编辑  ~/.profile
export PATH="$HOME/.cargo/bin:$PATH"
# 使配置生效
source ~/.profile
# 确认
rustc --version
cargo --version

## 注：这样配置每次使用都要重新 source 让 PATH 生效

# Set the default rustup version to 1.50.0 or lower
rustup default 1.50.0
# Install the rustwasmc
curl https://raw.githubusercontent.com/second-state/rustwasmc/master/installer/init.sh -sSf | sh
```