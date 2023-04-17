- create tmux session : tmux
- attach tmux: tmux a
- create new session: tmux new -s s_name
- attach session with session name: tmux a -t s_name
- list tmux session: tmux ls
- kill tmux session: tmux kill-session -t s_name


====


Ctrl+b c  # Tạo một cửa sổ mới
Ctrl+b w  # Xem danh sách cửa sổ hiện tại
Ctrl+b n/p  #　Chuyển đến cửa sổ tiếp theo hoặc trước đó
Ctrl+b f  #　Tìm kiếm cửa sổ
Ctrl+b ,  #　Đặt/Đổi tên cho cửa sổ
Ctrl+b &  #　Đóng cửa sổ

Ctrl+b %  # chia đôi màn hình theo chiều dọc
Ctrl+b "  # chia đôi màn hình theo chiều ngang

Ctrl+b o/<phím mũi tên>  # Di chuyển giữa các panel
Ctrl+b q  # Hiện số thứ tự trên
Ctrl+b x  # Xoá panel


- exmple to change ctrt+b to ctrl+c

tmux source-file .tmux.conf

unbind C-b
set -g prefix C-a
