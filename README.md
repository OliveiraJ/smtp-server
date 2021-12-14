# smtp-server

<p>This is a small application that uses the <a href=https://kb.synology.com/en-global/SRM/tutorial/How_to_use_Gmail_SMTP_server_to_send_emails_for_SRM>Gmail SMTP server</a>, built in Go. </p>
<p>Built as part of the ativities of the Computer Netorks course, developed only for learning purpose this program requires you to insert your Gmail and its Password (have in mind that none of those is stored or serve ny other purpose than authenticate your identidy so the email can be sent trough your gmail accout).</p>
<p><h3>How to use it</h3>
  <ul>
    <li>First of all, you should have the the option "Access to less secure apps" enabled in your google account (if you have 2FA enabled you will need to disabel it before).</li>
  <l1>Next step is to have your port 3000 free, the server runs in this port so make sure other applications is not using it while you are ruuning the code.</li>
  <li>This code was developed on Go v.1.17.3, so its recommedd that you have at least this version if you are looking to run or test this code.</li>
  <li>To use it you can open your terminal on the folder and run "go run main.go" or build the app usign "go build" generating the executable file.</li> 
  <li>With the app runnig accest the address localhost:3000 on your browser.</li>
  </ul>
</p>
