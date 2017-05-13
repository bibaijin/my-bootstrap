filename=清华大学_电子工程_张凯

main.pdf: main.tex
	latexmk -xelatex main

en: main-en.tex
	latexmk -xelatex main-en

topng:
	convert -density 300 main.pdf img/CV.png

rename:
	mv main.pdf "$(filename).pdf"

view:
	okular main.pdf &

clean:
	rm -f *.aux *.bbl *.blg *.idx *.ilg *.ind *.out *.toc *.nav *.snm main.pdf *.run.xml *-blx.bib *.bcf *.log *.fls *.fdb_latexmk
