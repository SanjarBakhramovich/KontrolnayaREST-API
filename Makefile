commit:
	@read -p "Введите текст для коммита: " message; \
	git add .; \
	git commit -m "$$message"; \
	git push;
