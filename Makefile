NAME = computerV1
SRCS = computerV1.go

all: $(NAME)

$(NAME):
	go build $(SRCS)

clean:
	

fclean: clean
	rm -rf $(NAME)

re: fclean all

.PHONY: all clean fclean re
