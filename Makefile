dist:
	rm -rf bin && \
	go build -o main && \
	mkdir bin && \
        mv main bin/
