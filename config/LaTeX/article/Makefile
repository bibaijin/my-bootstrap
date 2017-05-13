filename=周报-2016年-6月-23日

main.pdf: main.tex
	latexmk -xelatex main.tex

save:
	mv main.pdf "$(filename).pdf"

view:
	evince main.pdf &

clean:
	rm -f *.aux *.bbl *.blg *.idx *.ilg *.ind *.out *.toc *.nav *.snm main.pdf
	rm -f *.run.xml *-blx.bib *.bcf *.log *.fdb_latexmk *.fls
