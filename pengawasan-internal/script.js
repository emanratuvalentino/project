const $=s=>document.querySelector(s);
const $$=s=>document.querySelectorAll(s);

// Mobile navigation
$('#menuToggle').addEventListener('click',()=>$('#mobileNav').classList.toggle('show'));
$$('.mobile-nav a').forEach(a=>a.addEventListener('click',()=>$('#mobileNav').classList.remove('show')));

// Theme preference
const savedTheme=localStorage.getItem('pi-theme');
if(savedTheme==='dark') document.body.classList.add('dark');
$('#themeToggle').textContent=document.body.classList.contains('dark')?'☀':'☾';
$('#themeToggle').addEventListener('click',()=>{
  document.body.classList.toggle('dark');
  const dark=document.body.classList.contains('dark');
  localStorage.setItem('pi-theme',dark?'dark':'light');
  $('#themeToggle').textContent=dark?'☀':'☾';
});

// Search chapters
$('#searchInput').addEventListener('input',e=>{
  const q=e.target.value.trim().toLowerCase(); let visible=0;
  $$('.chapter-card').forEach(card=>{
    const ok=card.dataset.search.includes(q);
    card.style.display=ok?'flex':'none'; if(ok) visible++;
  });
  $('#noResult').hidden=visible!==0;
});

// Reading progress
addEventListener('scroll',()=>{
  const max=document.documentElement.scrollHeight-innerHeight;
  $('#progress').style.width=`${max?scrollY/max*100:0}%`;
},{passive:true});
$('#backTop').addEventListener('click',()=>scrollTo({top:0,behavior:'smooth'}));

// Quiz
const quiz=[
 {q:'Apa empat tujuan pengawasan internal?',o:['Efektivitas/efisiensi, keandalan laporan keuangan, pengamanan aset, ketaatan aturan','Menaikkan anggaran, menambah pegawai, membuat laporan, membeli aset','Audit, reviu, evaluasi, konsultansi'],a:0},
 {q:'Manakah yang termasuk unsur Sistem Pengendalian Intern?',o:['Promosi dan mutasi','Penilaian risiko','Pemasaran layanan'],a:1},
 {q:'Tahap apa yang dilakukan setelah pelaksanaan pengawasan?',o:['Perencanaan','Pelaporan','Rekrutmen'],a:1},
 {q:'Salah satu prinsip kode etik auditor adalah...',o:['Objektivitas','Kompetisi antar-OPD','Kerahasiaan publik tanpa batas'],a:0}
];
let qi=0,answered=false;
function renderQuiz(){
  const x=quiz[qi]; answered=false; $('#quizQuestion').textContent=x.q; $('#quizFeedback').textContent='';
  $('#quizNext').textContent=qi===quiz.length-1?'Selesai':'Pertanyaan berikutnya';
  $('#quizOptions').innerHTML=x.o.map((v,i)=>`<button class="quiz-option" data-i="${i}">${v}</button>`).join('');
  $$('.quiz-option').forEach(btn=>btn.addEventListener('click',()=>{
    if(answered)return; answered=true; const i=Number(btn.dataset.i);
    $$('.quiz-option').forEach((b,j)=>{if(j===x.a)b.classList.add('correct')});
    if(i===x.a) $('#quizFeedback').textContent='Benar!'; else {btn.classList.add('wrong');$('#quizFeedback').textContent='Belum tepat. Jawaban yang sesuai ditandai.'}
  }));
}
function openQuiz(){qi=0;renderQuiz();$('#quizModal').classList.add('show');$('#quizModal').setAttribute('aria-hidden','false')}
function closeQuiz(){$('#quizModal').classList.remove('show');$('#quizModal').setAttribute('aria-hidden','true')}
$('#quizOpen').addEventListener('click',openQuiz);$('#quizClose').addEventListener('click',closeQuiz);
$('#quizModal').addEventListener('click',e=>{if(e.target.id==='quizModal')closeQuiz()});
$('#quizNext').addEventListener('click',()=>{if(!answered){$('#quizFeedback').textContent='Pilih salah satu jawaban dulu.';return}if(qi<quiz.length-1){qi++;renderQuiz()}else{closeQuiz();alert('Selesai! Terima kasih sudah mengecek pemahaman materi.')}});
