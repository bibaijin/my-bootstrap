filename=答辩PPT-2016-5-25-张凯

main.pdf:main.tex
	latexmk -xelatex main.tex

save:
	cp main.pdf "$(filename).pdf"

view:
	okular main.pdf &

clean:
	rm -f *.aux *.bbl *.blg *.idx *.ilg *.ind *.out *.toc *.nav *.snm main.pdf *.run.xml *-blx.bib *.bcf *.log *.fls *.fdb_latexmk
