#include<iostream>
#include<string>

using namespace std;

class Solution{
public:
	bool isInterleave(string s1, string s2, string s3) {
		int width = s1.length();
		int length = s2.length();

		bool result = false;
		
		bool ** bi_mat = new bool*[length+1];
		
		for (int i = 0; i < length+1; ++i) {
			bi_mat[i] = new bool[width+1];
			for (int j = 0; j < width + 1; ++j) {
				bi_mat[i][j] =false;
			}
		}
		
		if (width + length != s3.length())
			return false;
		bi_mat[0][0] = true;
		for (int i = 0; i < length + 1; ++i) {
			for (int j = 0; j < width + 1; ++j) {
				if (j > 0&&s1.at(j-1) == s3.at(i+j-1)) {
					bi_mat[i][j] |= bi_mat[i][j-1];
				}
				if (i > 0 && s2.at(i-1) == s3.at(i+j-1)) {
					bi_mat[i][j] |= bi_mat[i-1][j];
				}
			}
		}
		
		result = bi_mat[length][width];
		
		for (int i = 0; i < length+1; ++i) {
			delete[] bi_mat[i];
		}
		delete[] bi_mat;
		
		return result;
	}
};

int main()
{
	string s1 = "aabcc";
	string s2 = "";
	Solution s;
	cout << s.isInterleave(s1, s2, "aadbb") << endl;
	cout << s.isInterleave(s1, s2, "aadbbbaccc") << endl;
}

