package org.edgegallery.example_app.util;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;
import java.util.StringTokenizer;

import org.apache.commons.lang.StringUtils;
import org.edgegallery.example_app.model.EALTEdgeBackup;
import org.edgegallery.example_app.model.EALTEdgeRestore;
import org.springframework.stereotype.Service;


@Service
public class ShellCommand {

	 public String executeCommand(String command) {

	    StringBuffer output = new StringBuffer();

	    Process p;
	    try {

	        p = Runtime.getRuntime().exec(command);
	        p.waitFor();
			BufferedReader reader = new BufferedReader(new InputStreamReader(p.getInputStream()));

	        String line = "";
	        while ((line = reader.readLine())!= null) {
	            output.append(line + "\n");
	        }
		}
		catch (Exception e) {
	        e.printStackTrace();
	    }
	    return output.toString();

	}

    public List<EALTEdgeBackup> executeBackupCommand(String command) {

        EALTEdgeBackup backup = new EALTEdgeBackup();
		List<EALTEdgeBackup> backupsList = new ArrayList<EALTEdgeBackup>();

		try {
			Process p;
            p = Runtime.getRuntime().exec(command);
            p.waitFor();
            BufferedReader reader =
                    new BufferedReader(new InputStreamReader(p.getInputStream()));

            String line = "";
            while ((line = reader.readLine())!= null) {
            	if(line.startsWith("NAME")) {
            		continue;
            	}
            	else {
            		backup = parseBackupResult(line);
                	backupsList.add(backup);
            	}
            }

        } catch (Exception e) {
            e.printStackTrace();
        }
        return backupsList;
    }
    
    public static EALTEdgeBackup parseBackupResult(String newstr){

    	EALTEdgeBackup backup = new EALTEdgeBackup();
    	List<String> al = new ArrayList<String>();
		
		StringTokenizer st = new StringTokenizer(newstr, " ");
		StringBuffer sb = new StringBuffer();
		
		while(st.hasMoreElements()) {
			sb.append(st.nextElement()).append(" ");
		}
		
		String newstrwithProperSpacing = sb.toString();
		String str[] = newstrwithProperSpacing.split(" ");
		
		str[4] = str[4] + str[5] + str[6] + str[7];
		
		al = Arrays.asList(str);
		
		for(int i = 0; i < al.size(); i++) {
			if( i == 0 ) {
				backup.setName(al.get(i));
			}
			if( i == 1) {
				backup.setStatus(al.get(i));
			}
			if( i == 2) {
				backup.setErrors(al.get(i));
			}
			if( i == 3){
				backup.setWarnings(al.get(i));
			}
			if( i == 4) {
				backup.setCreated(al.get(i));
			}

		}
        return backup;
    }

    public List<EALTEdgeRestore> executeRestoreCommand(String command) {

        EALTEdgeRestore restore = new EALTEdgeRestore();
        List<EALTEdgeRestore> restoresList = new ArrayList<EALTEdgeRestore>();
        
        try {
        	Process p;
            p = Runtime.getRuntime().exec(command);
            p.waitFor();
            BufferedReader reader =
                    new BufferedReader(new InputStreamReader(p.getInputStream()));

            String line = "";
            while ((line = reader.readLine())!= null) {
            	if(line.startsWith("NAME")) {
            		continue;
            	}
            	else {
            		restore = parseRestoreResult(line);
                	restoresList.add(restore);
            	}
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
        return restoresList;
    }

    public static EALTEdgeRestore parseRestoreResult(String newstr){

    	EALTEdgeRestore restore = new EALTEdgeRestore();

    	StringTokenizer st = new StringTokenizer(newstr, " ");
		StringBuffer sb = new StringBuffer();

		while(st.hasMoreElements()) {
			sb.append(st.nextElement()).append(" ");
		}

		String newstrwithProperSpacing = sb.toString();
		String str[] = newstrwithProperSpacing.split(" ");
		
		List<String> ll = new LinkedList<String>(Arrays.asList(str));
		
		for(int i = 0; i < ll.size(); i++) {
			if( i == 0 ) {
				restore.setName(ll.get(i));
			}
			if( i == 1) {
				restore.setBackup(ll.get(i));
			}
			if( i == 2) {
				restore.setStatus(ll.get(i));
			}
		}

        return restore;
    }
}
